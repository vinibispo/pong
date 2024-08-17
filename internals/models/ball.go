package models

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	X      float32
	Y      float32
	SpeedX int
	SpeedY int
	Radius int
}

func NewBall(x, y float32, speedX, speedY, radius int) *Ball {
	return &Ball{x, y, speedX, speedY, radius}
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), float32(b.Radius), rl.Red)
}

func (b *Ball) Update() {
	b.X += float32(b.SpeedX)
	b.Y += float32(b.SpeedY)

	if b.Y+float32(b.Radius) >= float32(rl.GetScreenHeight()) || b.Y-float32(b.Radius) <= 0 {
		b.SpeedY = -b.SpeedY
	}

	if b.X+float32(b.Radius) >= float32(rl.GetScreenWidth()) || b.X-float32(b.Radius) <= 0 {
		b.SpeedX = -b.SpeedX
	}
}
