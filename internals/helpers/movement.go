package helpers

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddler interface {
	//Fix add Y and Height fields

	GetY() float32
	GetHeight() float32
	UpdateY(float32)
}

func LimitMovement(p Paddler) {
	y := p.GetY()
	height := p.GetHeight()
	if y <= 0 {
		p.UpdateY(0)
	}

	if y+height >= float32(rl.GetScreenHeight()) {
		p.UpdateY(float32(rl.GetScreenHeight()) - height)
	}
}
