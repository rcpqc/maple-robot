package ix

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"log"
	"os/exec"
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
	frameBuffer *image.RGBA
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

	o.frameBuffer = &image.RGBA{
		Pix:    buf.Bytes()[:width*4*height],
		Stride: 4 * int(width),
		Rect:   image.Rect(0, 0, int(width), int(height)),
	}

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

func GetPixel(pos Position) Color {
	c := Display.frameBuffer.RGBAAt(int(pos.X), int(pos.Y))
	return Color{R: c.R, G: c.G, B: c.B}
}

func WaitPixel(pos Position, c Color) {
	for {
		if GetPixel(pos).Equals(c) {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
