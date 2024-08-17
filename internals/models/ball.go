package models

import (
	"pong/internals/helpers"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	rl.DrawCircle(int32(b.X), int32(b.Y), float32(b.Radius), helpers.Yellow)
}

func (b *Ball) Update(playerScore, cpuScore *int) {
	b.X += float32(b.SpeedX)
	b.Y += float32(b.SpeedY)

	if b.Y+float32(b.Radius) >= float32(rl.GetScreenHeight()) || b.Y-float32(b.Radius) <= 0 {
		b.SpeedY = -b.SpeedY
	}

	if b.X+float32(b.Radius) >= float32(rl.GetScreenWidth()) {
		*cpuScore++
		b.Reset()
	}

	if b.X-float32(b.Radius) <= 0 {
		*playerScore++
		b.Reset()
	}
}

func (b *Ball) Reset() {
	b.X = float32(rl.GetScreenWidth()) / 2
	b.Y = float32(rl.GetScreenHeight()) / 2

	speedChoices := []int{1, -1}
	b.SpeedX *= speedChoices[rl.GetRandomValue(0, 1)]
	b.SpeedY *= speedChoices[rl.GetRandomValue(0, 1)]
}
