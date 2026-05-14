package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"sync"

	"maple-robot/ix"
)

//go:embed static/index.html
var staticFiles embed.FS

// LogHub 接收日志消息并缓存最近 N 条, 供 HTTP 轮询.
type LogHub struct {
	mu     sync.Mutex
	buffer []string // 环形缓冲区
	idx    int      // 当前写入位置
	cap    int      // 最大条数
	full   bool     // 是否已填满
}

func NewLogHub() *LogHub {
	return &LogHub{
		buffer: make([]string, 500),
		cap:    500,
	}
}

func (h *LogHub) Write(p []byte) (int, error) {
	msg := string(p)
	h.mu.Lock()
	h.buffer[h.idx] = msg
	h.idx++
	if h.idx >= h.cap {
		h.idx = 0
		h.full = true
	}
	h.mu.Unlock()
	return len(p), nil
}

// RecentLogs 返回最近的 n 条日志.
func (h *LogHub) RecentLogs(n int) []string {
	h.mu.Lock()
	defer h.mu.Unlock()
	total := h.idx
	if h.full {
		total = h.cap
	}
	if n > total {
		n = total
	}
	result := make([]string, n)
	for i := 0; i < n; i++ {
		pos := (h.idx - n + i) % h.cap
		if pos < 0 {
			pos += h.cap
		}
		result[i] = h.buffer[pos]
	}
	return result
}

// Start launches the web dashboard server (blocking).
func Start(addr string, logHub *LogHub, runner *Runner) error {
	mux := http.NewServeMux()

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		return fmt.Errorf("fs sub -> %w", err)
	}
	mux.Handle("/", http.FileServer(http.FS(staticFS)))

	mux.HandleFunc("/screenshot", handleScreenshot)
	mux.HandleFunc("/api/tap", handleTap)
	mux.HandleFunc("/api/swipe", handleSwipe)
	mux.HandleFunc("/api/info", handleInfo)
	mux.HandleFunc("/api/start", func(w http.ResponseWriter, r *http.Request) {
		handleStart(w, r, runner)
	})
	mux.HandleFunc("/api/stop", func(w http.ResponseWriter, r *http.Request) {
		handleStop(w, r, runner)
	})
	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, runner)
	})
	mux.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		handleLogs(w, r, logHub)
	})

	log.Printf("[web] dashboard listening on %s", addr)
	return http.ListenAndServe(addr, mux)
}

// ---- screenshot ----

func handleScreenshot(w http.ResponseWriter, r *http.Request) {
	data, wd, ht, err := ix.CaptureJPEG()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("X-Display-Width", fmt.Sprintf("%d", wd))
	w.Header().Set("X-Display-Height", fmt.Sprintf("%d", ht))
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Write(data)
}

// ---- tap / swipe ----

type tapReq struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

func handleTap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	var req tapReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ix.Tap(ix.Position{X: req.X, Y: req.Y})
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}

type swipeReq struct {
	X1       int64 `json:"x1"`
	Y1       int64 `json:"y1"`
	X2       int64 `json:"x2"`
	Y2       int64 `json:"y2"`
	Duration int64 `json:"duration"`
}

func handleSwipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	var req swipeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Duration <= 0 {
		req.Duration = 200
	}
	ix.Swipe(ix.Position{X: req.X1, Y: req.Y1}, ix.Position{X: req.X2, Y: req.Y2}, req.Duration)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}

// ---- info ----

type infoResp struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	wd, ht := ix.Display.Size()
	resp := infoResp{Width: wd, Height: ht}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ---- runner control ----

func handleStart(w http.ResponseWriter, r *http.Request, runner *Runner) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	if err := runner.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "started"})
}

func handleStop(w http.ResponseWriter, r *http.Request, runner *Runner) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	runner.Stop()
	json.NewEncoder(w).Encode(map[string]string{"status": "stopped"})
}

func handleStatus(w http.ResponseWriter, r *http.Request, runner *Runner) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(runner.Status())
}

// ---- log polling ----

func handleLogs(w http.ResponseWriter, r *http.Request, hub *LogHub) {
	logs := hub.RecentLogs(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
