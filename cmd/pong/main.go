package main

import (
	"fmt"
	"pong/internals/helpers"
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

	playerScore := 0
	cpuScore := 0
	ball := models.NewBall(screenWidth/2, screenHeight/2, 7, 7, ballRadius)
	playerWidth := 25
	playerHeight := 120
	player1 := models.NewPaddle(float32(screenWidth-playerWidth-10), float32((screenHeight-playerHeight)/2), float32(playerWidth), float32(playerHeight), 6)
	cpu := models.NewCPUPaddle(10, float32((screenHeight-playerHeight)/2), float32(playerWidth), float32(playerHeight), 6)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		ball.Update(&playerScore, &cpuScore)
		player1.Update()
		cpu.Update(int(ball.Y))

		if rl.CheckCollisionCircleRec(rl.Vector2{X: ball.X, Y: ball.Y}, float32(ball.Radius), rl.Rectangle{X: player1.X, Y: player1.Y, Width: player1.Width, Height: player1.Height}) {
			ball.SpeedX = -ball.SpeedX
		}

		if rl.CheckCollisionCircleRec(rl.Vector2{X: ball.X, Y: ball.Y}, float32(ball.Radius), rl.Rectangle{X: cpu.X, Y: cpu.Y, Width: cpu.Width, Height: cpu.Height}) {
			ball.SpeedX = -ball.SpeedX
		}

		rl.ClearBackground(helpers.DarkGreen)
		rl.DrawRectangle(screenWidth/2, 0, screenWidth/2, screenHeight, helpers.Green)
		rl.DrawCircle(screenWidth/2, screenHeight/2, 150, helpers.LightGreen)
		rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.White)
		ball.Draw()
		cpu.Draw()
		player1.Draw()
		rl.DrawText(fmt.Sprint(cpuScore), screenWidth/4-20, 20, 80, rl.White)
		rl.DrawText(fmt.Sprint(playerScore), 3*screenWidth/4-20, 20, 80, rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
