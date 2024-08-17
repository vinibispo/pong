package models

import rl "github.com/gen2brain/raylib-go/raylib"

type CPUPaddle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Speed  int32
}

func NewCPUPaddle(x, y, width, height float32, speed int32) *CPUPaddle {
	return &CPUPaddle{x, y, width, height, speed}
}

func (p *CPUPaddle) Draw() {
	rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), rl.White)
}

func (p *CPUPaddle) Update(ballY int) {
	if (p.Y + p.Height/2) > float32(ballY) {
		p.Y -= float32(p.Speed)
	}

	if (p.Y + p.Height/2) < float32(ballY) {
		p.Y += float32(p.Speed)
	}

	if p.Y <= 0 {
		p.Y = 0
	}

	if p.Y+p.Height >= float32(rl.GetScreenHeight()) {
		p.Y = float32(rl.GetScreenHeight()) - p.Height
	}
}
