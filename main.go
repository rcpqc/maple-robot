package main

import (
	"context"
	"fmt"
	"os"
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
	// Web 日志广播 hub (运行时按当天追加日志文件)
	logHub := web.NewLogHub()

	// 读取服务配置
	cfg, err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	// 创建 Runner (点击"开始"时按当天加载角色与日志, 支持跨天启动)
	runner := web.NewRunner(logHub)

	// 服务地址
	addr := cfg.Web.Addr
	if addr == "" {
		addr = ":8080"
	}
	localAddr := "127.0.0.1" + addr

	fmt.Printf("=== Maple Robot 服务已启动 ===\n")
	fmt.Printf("    Web 面板: http://%s\n", localAddr)
	fmt.Printf("    按回车键开始挂机\n")
	fmt.Println()

	// 监听回车键启动挂机
	go func() {
		buf := make([]byte, 1)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				return
			}
			if n > 0 && (buf[0] == '\n' || buf[0] == '\r') {
				if err := runner.Start(); err != nil {
					fmt.Printf("[enter] 启动失败: %v\n", err)
				} else {
					fmt.Printf("[enter] 挂机已启动\n")
				}
				return
			}
		}
	}()

	// 启动 Adele 隧道 (非阻塞)
	if cfg.Adele != nil && cfg.Adele.Server != "" {
		go startAdele(cfg.Adele, localAddr)
	}

	fmt.Println()

	// 启动 Web 服务 (阻塞)
	if err := web.Start(addr, logHub, runner, cfg.Auth); err != nil {
		log.Warn(context.Background(), "web server stopped", "err", err)
	}
}

func startAdele(cfg *config.AdeleConfig, localAddr string) {
	clientID := cfg.ClientID
	if clientID == "" {
		clientID = "maple-robot"
	}

	for {
		c := client.New(clientID, localAddr, client.WithServerAddress(cfg.Server))

		fmt.Printf("[Adele] 连接中 %s ...\n", cfg.Server)
		if err := c.Connect(); err != nil {
			fmt.Printf("[Adele] 连接失败: %v, 5秒后重试\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("[Adele] 隧道已建立: %s → %s\n", localAddr, c.ProxyAddr())

		// 等待 gRPC stream 断开 (recvLoop 退出即触发重连)
		c.Wait()

		fmt.Printf("[Adele] 隧道断开, 3秒后重连...\n")
		c.Disconnect()
		time.Sleep(3 * time.Second)
	}
}
