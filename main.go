package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	cellSize = 20
	gridW    = 30
	gridH    = 24
)

type Point struct {
	X, Y int
}

type Game struct {
	snake     []Point
	direction Point
	tick      int
	speed     int
}

func (g *Game) Update() error {
	// --- INPUT ---
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.direction.Y != 1 {
		g.direction = Point{0, -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.direction.Y != -1 {
		g.direction = Point{0, 1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.direction.X != 1 {
		g.direction = Point{-1, 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.direction.X != -1 {
		g.direction = Point{1, 0}
	}

	// --- TICK CONTROL ---
	g.tick++
	if g.tick%g.speed != 0 {
		return nil
	}

	// --- MOVE SNAKE ---
	head := g.snake[0]
	newHead := Point{
		X: head.X + g.direction.X,
		Y: head.Y + g.direction.Y,
	}

	// wrap around screen (no death yet)
	if newHead.X < 0 {
		newHead.X = gridW - 1
	}
	if newHead.X >= gridW {
		newHead.X = 0
	}
	if newHead.Y < 0 {
		newHead.Y = gridH - 1
	}
	if newHead.Y >= gridH {
		newHead.Y = 0
	}

	g.snake = append([]Point{newHead}, g.snake[:len(g.snake)-1]...)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.snake {
		vector.FillRect(
			screen,
			float32(p.X*cellSize),
			float32(p.Y*cellSize),
			cellSize,
			cellSize,
			color.RGBA{0, 255, 0, 255},
			false,
		)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return gridW * cellSize, gridH * cellSize
}

func main() {
	game := &Game{
		snake: []Point{
			{5, 5},
			{4, 5},
			{3, 5},
		},
		direction: Point{1, 0}, // start moving right
		speed:     6,           // lower = faster
	}

	ebiten.SetWindowTitle("Go Snake 🐍")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
