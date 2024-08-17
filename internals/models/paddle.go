package models

import rl "github.com/gen2brain/raylib-go/raylib"
import helpers "pong/internals/helpers"

type Paddle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Speed  int32
}

func NewPaddle(x, y, width, height float32, speed int32) *Paddle {
	return &Paddle{x, y, width, height, speed}
}

func (p *Paddle) Draw() {
	rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), rl.White)
}

func (p *Paddle) GetY() float32 {
	return float32(p.Y)
}

func (p *Paddle) GetHeight() float32 {
	return float32(p.Height)
}

func (p *Paddle) UpdateY(y float32) {
	p.Y = y
}

func (p *Paddle) Update() {
	if rl.IsKeyDown(rl.KeyUp) {
		p.Y -= float32(p.Speed)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.Y += float32(p.Speed)
	}

	helpers.LimitMovement(p)
}
