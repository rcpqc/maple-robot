package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"maple-robot/config"
	"maple-robot/ix"
	"maple-robot/log"
	"maple-robot/web"

	"github.com/rcpqc/adele/client"
)

func init() {
	if err := ix.Check(); err != nil {
		panic(err)
	}
}

func main() {
	// 日志文件 (按天)
	logFile := "logs/" + time.Now().Format(time.DateOnly) + ".log"
	records, _ := config.LoadTaskRecords(logFile)
	ctx := config.WithRecords(context.Background(), records)

	// 设置日志输出: stdout + 文件 + Web 广播
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	logHub := web.NewLogHub()
	baseLogger := log.New(io.MultiWriter(os.Stdout, f, logHub))

	// 读取服务配置
	cfg, err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	// 读取角色 & 脚本
	roles, err := config.LoadRoles()
	if err != nil {
		panic(err)
	}

	// 创建 Runner (后续通过 Web 面板启动/停止)
	runner := web.NewRunner(roles, baseLogger, logFile)

	// 服务地址
	addr := cfg.Web.Addr
	if addr == "" {
		addr = ":8080"
	}
	localAddr := "127.0.0.1" + addr

	fmt.Printf("=== Maple Robot 服务已启动 ===\n")
	fmt.Printf("    Web 面板: http://%s\n", localAddr)
	fmt.Printf("    日志文件: %s\n", logFile)
	fmt.Printf("    角色数量: %d\n", len(roles))

	// 启动 Adele 隧道 (非阻塞)
	if cfg.Adele != nil && cfg.Adele.Server != "" {
		go startAdele(cfg.Adele, localAddr)
	}

	fmt.Println()

	// 启动 Web 服务 (阻塞)
	if err := web.Start(addr, logHub, runner, cfg.Auth); err != nil {
		log.Warn(ctx, "web server stopped", "err", err)
	}
}

func startAdele(cfg *config.AdeleConfig, localAddr string) {
	clientID := cfg.ClientID
	if clientID == "" {
		clientID = "maple-robot"
	}

	serverHost, _, _ := net.SplitHostPort(cfg.Server)

	for {
		c := client.New(clientID, localAddr, client.WithServerAddress(cfg.Server))

		fmt.Printf("[Adele] 连接中 %s ...\n", cfg.Server)
		if err := c.Connect(); err != nil {
			fmt.Printf("[Adele] 连接失败: %v, 5秒后重试\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		proxyAddr := c.ProxyAddr()
		proxyPort := strings.TrimPrefix(proxyAddr, ":")
		healthURL := fmt.Sprintf("https://%s:%s/api/info", serverHost, proxyPort)

		fmt.Printf("[Adele] 隧道已建立: %s → https://%s:%s\n", localAddr, serverHost, proxyPort)

		// 健康检查: 每15秒通过隧道发请求, 连续2次失败则重连
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		hc := &http.Client{Timeout: 10 * time.Second, Transport: tr}

		failures := 0
		ticker := time.NewTicker(15 * time.Second)
		for range ticker.C {
			resp, err := hc.Get(healthURL)
			if err != nil {
				failures++
				fmt.Printf("[Adele] 健康检查失败 (%d/2): %v\n", failures, err)
			} else {
				resp.Body.Close()
				failures = 0
			}
			if failures >= 2 {
				break
			}
		}
		ticker.Stop()

		fmt.Printf("[Adele] 隧道断开, 3秒后重连...\n")
		c.Disconnect()
		time.Sleep(3 * time.Second)
	}
}
