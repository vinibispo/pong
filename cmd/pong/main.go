package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1280
	screenHeight = 800
	ballRadius   = 20
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

		rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.White)
		rl.DrawCircle(screenWidth/2, screenHeight/2, ballRadius, rl.Maroon)
		rl.DrawRectangle(10, screenHeight/2-60, 25, 120, rl.White)
		rl.DrawRectangle(screenWidth-35, screenHeight/2-60, 25, 120, rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
