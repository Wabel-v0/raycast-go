package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var gridMap = [10][10]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 1, 0, 0, 0, 1, 0, 1},
	{1, 1, 0, 0, 1, 0, 0, 0, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
	{1, 0, 1, 0, 1, 0, 0, 1, 1, 1},
	{1, 1, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 0, 0, 1, 0, 0, 1, 0, 1},
	{1, 0, 1, 0, 0, 0, 1, 0, 1, 1},
	{1, 0, 1, 0, 1, 1, 0, 1, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

const tile = 40

type Player struct {
	px float32
	py float32
}

func (p Player) render2dPlayer() {
	rl.DrawCircle(int32(p.px*tile), int32(p.py*tile), 5, rl.Blue)

}
func (p *Player) movePlayer() {
	if rl.IsKeyDown(rl.KeyW) {
		p.py -= 0.1
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.px -= 0.1
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.py += 0.1
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.px += 0.1
	}
}
func render2dMap(gridMap [10][10]int) {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			color := rl.Brown
			if gridMap[row][col] == 1 {
				color = rl.Black
			}
			rl.DrawRectangle(int32(row*tile), int32(col*tile), tile, tile, color)
			rl.DrawRectangleLines(int32(row*tile), int32(col*tile), tile, tile, rl.White)

		}
	}
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	player := Player{
		px: 2.5,
		py: 2.5,
	}
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		render2dMap(gridMap)
		player.render2dPlayer()
		player.movePlayer()
		rl.EndDrawing()
	}
}
