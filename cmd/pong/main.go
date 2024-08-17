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
	playerWidth := 25
	playerHeight := 120
	player1 := models.NewPaddle(float32(screenWidth-playerWidth-10), float32((screenHeight-playerHeight)/2), float32(playerWidth), float32(playerHeight), 6)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		ball.Update()
		player1.Update()

		rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.White)
		ball.Draw()
		rl.DrawRectangle(10, screenHeight/2-60, 25, 120, rl.White)
		player1.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
