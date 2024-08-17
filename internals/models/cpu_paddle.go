package models

import (
	"pong/internals/helpers"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CPUPaddle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Speed  int32
}

func (p *CPUPaddle) GetHeight() float32 {
	return p.Height
}

func (p *CPUPaddle) GetY() float32 {
	return p.Y
}

func (p *CPUPaddle) UpdateY(y float32) {
	p.Y = y
}

func NewCPUPaddle(x, y, width, height float32, speed int32) *CPUPaddle {
	return &CPUPaddle{x, y, width, height, speed}
}

func (p *CPUPaddle) Draw() {
	rect := rl.Rectangle{X: p.X, Y: p.Y, Width: p.Width, Height: p.Height}
	rl.DrawRectangleRounded(rect, 0.8, 0, rl.White)
}

func (p *CPUPaddle) Update(ballY int) {
	if (p.Y + p.Height/2) > float32(ballY) {
		p.Y -= float32(p.Speed)
	}

	if (p.Y + p.Height/2) < float32(ballY) {
		p.Y += float32(p.Speed)
	}

	helpers.LimitMovement(p)
}
