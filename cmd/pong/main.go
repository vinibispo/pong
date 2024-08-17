package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 600
	ballRadius   = 15
)

func main() {
	ballX := int32(100)
	ballY := int32(100)
	ballSpeedX := 5
	ballSpeedY := 5
	rl.InitWindow(screenWidth, screenHeight, "My pong game!")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		ballX += int32(ballSpeedX)
		ballY += int32(ballSpeedY)

		if ballX >= screenWidth-ballRadius || ballX <= ballRadius {
			ballSpeedX *= -1
		}

		if ballY >= screenHeight-ballRadius || ballY <= ballRadius {
			ballSpeedY *= -1
		}

		rl.DrawCircle(ballX, ballY, ballRadius, rl.Maroon)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
