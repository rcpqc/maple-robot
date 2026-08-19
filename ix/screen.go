package ix

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"math"
	"os/exec"
	"sync"
	"time"
)

var Display *Screen = &Screen{}

func Check() error {
	if err := Display.take(); err != nil {
		return fmt.Errorf("screen take -> %w", err)
	}
	Display.update(time.Second)
	return nil
}

type Screen struct {
	mu          sync.RWMutex
	frameBuffer *image.RGBA
	jpegCache   []byte // 预编码的 JPEG 缓存 (缩小版)
	jpegW       int    // 缓存图片宽度
	jpegH       int    // 缓存图片高度
}

func (o *Screen) take() error {
	buf := bytes.NewBuffer(nil)
	cmd := exec.Command("adb", "exec-out", "screencap")
	// cmd := exec.Command("screencap")
	cmd.Stdout = buf

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run -> %v", err)
	}

	var width, height, format, reverse int32
	binary.Read(buf, binary.LittleEndian, &width)
	binary.Read(buf, binary.LittleEndian, &height)
	binary.Read(buf, binary.LittleEndian, &format)
	binary.Read(buf, binary.LittleEndian, &reverse)

	fb := &image.RGBA{
		Pix:    buf.Bytes()[:width*4*height],
		Stride: 4 * int(width),
		Rect:   image.Rect(0, 0, int(width), int(height)),
	}

	// 编码 JPEG (质量 60)
	enc := new(bytes.Buffer)
	if err := jpeg.Encode(enc, fb, &jpeg.Options{Quality: 60}); err != nil {
		return fmt.Errorf("jpeg cache -> %w", err)
	}

	o.mu.Lock()
	o.frameBuffer = fb
	o.jpegCache = enc.Bytes()
	o.jpegW = int(width)
	o.jpegH = int(height)
	o.mu.Unlock()

	return nil
}

func (o *Screen) update(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			if err := o.take(); err != nil {
				log.Printf("take -> %v\n", err)
			}
		}
	}()
}

func (o *Screen) Size() (int, int) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	if o.frameBuffer == nil {
		return 0, 0
	}
	return o.frameBuffer.Rect.Dx(), o.frameBuffer.Rect.Dy()
}

func GetPixel(pos Position) Color {
	Display.mu.RLock()
	defer Display.mu.RUnlock()
	if Display.frameBuffer == nil {
		return Color{}
	}
	c := Display.frameBuffer.RGBAAt(int(pos.X), int(pos.Y))
	return Color{R: c.R, G: c.G, B: c.B}
}

// FindPixelInColumn 在 x=col 这一列中，从上到下查找第一个与目标颜色 target
// 欧氏距离小于 threshold 的像素，返回其 Position 与 true；未找到返回零值与 false。
func FindPixelInColumn(col int64, target Color, threshold float64) (Position, bool) {
	Display.mu.RLock()
	defer Display.mu.RUnlock()
	if Display.frameBuffer == nil {
		return Position{}, false
	}
	w := Display.frameBuffer.Rect.Dx()
	h := Display.frameBuffer.Rect.Dy()
	if int(col) < 0 || int(col) >= w {
		return Position{}, false
	}
	for y := 0; y < h; y++ {
		c := Display.frameBuffer.RGBAAt(int(col), y)
		dr := float64(c.R) - float64(target.R)
		dg := float64(c.G) - float64(target.G)
		db := float64(c.B) - float64(target.B)
		dist := math.Sqrt(dr*dr + dg*dg + db*db)
		if dist < threshold {
			return Position{X: col, Y: int64(y)}, true
		}
	}
	return Position{}, false
}

func WaitPixel(pos Position, c Color) {
	for {
		if GetPixel(pos).Equals(c) {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func SubImage(rc Rect) image.Image {
	Display.mu.RLock()
	defer Display.mu.RUnlock()
	if Display.frameBuffer == nil {
		return nil
	}
	return Display.frameBuffer.SubImage(image.Rect(int(rc.X), int(rc.Y), int(rc.W), int(rc.H)))
}

// CaptureJPEG 返回预编码的 JPEG 缓存, 不会触发新的编码.
func CaptureJPEG() ([]byte, int, int, error) {
	Display.mu.RLock()
	cache := Display.jpegCache
	w, h := Display.jpegW, Display.jpegH
	Display.mu.RUnlock()
	if cache == nil {
		return nil, 0, 0, fmt.Errorf("jpeg cache not ready")
	}
	data := make([]byte, len(cache))
	copy(data, cache)
	return data, w, h, nil
}
