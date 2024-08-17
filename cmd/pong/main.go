package main

import (
	"pong/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenHeight = 800
	ballRadius   = 20
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "My pong game!")
	rl.SetTargetFPS(60)

	ball := models.NewBall(screenWidth/2, screenHeight/2, 7, 7, ballRadius)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		ball.Update()

		rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.White)
		ball.Draw()
		rl.DrawRectangle(10, screenHeight/2-60, 25, 120, rl.White)
		rl.DrawRectangle(screenWidth-35, screenHeight/2-60, 25, 120, rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
