package main

import rl "github.com/gen2brain/raylib-go/raylib"

var gridMap = [10][10]int{
	{1, 0, 1, 0, 0, 1, 1, 0, 1, 0},
	{0, 1, 0, 1, 1, 0, 0, 1, 0, 1},
	{1, 1, 0, 0, 1, 0, 1, 0, 1, 0},
	{0, 0, 1, 1, 0, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 1, 0, 0, 1, 1, 0},
	{0, 1, 0, 1, 0, 1, 1, 0, 0, 1},
	{1, 1, 0, 0, 1, 1, 0, 1, 0, 0},
	{0, 0, 1, 1, 0, 0, 1, 0, 1, 1},
	{1, 0, 1, 0, 1, 1, 0, 1, 0, 1},
	{0, 1, 0, 1, 0, 0, 1, 0, 1, 0},
}

const tile = 40

func render2dMap(gridMap [10][10]int) {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			color := rl.Brown
			if gridMap[row][col] == 1 {
				color = rl.Black
			}
			rl.DrawRectangle(int32(row*tile), int32(col*tile), tile, tile, color)

		}
	}
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		render2dMap(gridMap)
		rl.EndDrawing()
	}
}
