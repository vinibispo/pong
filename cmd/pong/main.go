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
	cpu := models.NewCPUPaddle(10, float32((screenHeight-playerHeight)/2), float32(playerWidth), float32(playerHeight), 6)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		ball.Update()
		player1.Update()
		cpu.Update(int(ball.Y))

		if rl.CheckCollisionCircleRec(rl.Vector2{ball.X, ball.Y}, float32(ball.Radius), rl.Rectangle{player1.X, player1.Y, player1.Width, player1.Height}) {
			ball.SpeedX = -ball.SpeedX
		}

		if rl.CheckCollisionCircleRec(rl.Vector2{ball.X, ball.Y}, float32(ball.Radius), rl.Rectangle{cpu.X, cpu.Y, cpu.Width, cpu.Height}) {
			ball.SpeedX = -ball.SpeedX
		}

		rl.ClearBackground(rl.Black)
		rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.White)
		ball.Draw()
		cpu.Draw()
		player1.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
