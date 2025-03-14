package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	CellSize     = 10
	UpdateSpeed  = 1 // lower value = faster updates
	ButtonWidth  = 120
	ButtonHeight = 30
)

var (
	ScreenWidth  = 1200
	ScreenHeight = 900
	Cols         = ScreenWidth / CellSize
	Rows         = ScreenHeight / CellSize
)

type Button struct {
	label  string
	x, y   int
	action func()
}

type Game struct {
	grid       [][]bool
	generation int
	tickCount  int
	mouseDown  bool
	running    bool
	buttons    []Button
}

func NewGame() *Game {
	g := &Game{
		grid:    make([][]bool, Rows),
		running: false,
	}
	for i := range g.grid {
		g.grid[i] = make([]bool, Cols)
	}
	g.buttons = []Button{
		{"Start", ScreenWidth/2 - 180, ScreenHeight - 50, g.ToggleRunning},
		{"Reset", ScreenWidth/2 - 60, ScreenHeight - 50, g.Reset},
		{"Random Seed", ScreenWidth/2 + 60, ScreenHeight - 50, g.Randomize},
	}
	return g
}

func (g *Game) Update() error {
	// handle mouse click and drag to toggle cells
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		col, row := x/CellSize, y/CellSize
		if row >= 0 && row < Rows && col >= 0 && col < Cols {
			g.grid[row][col] = !g.grid[row][col]
		}
	}

	// handle button clicks
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.mouseDown {
		g.mouseDown = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !g.mouseDown {
		g.mouseDown = true
		x, y := ebiten.CursorPosition()
		for _, btn := range g.buttons {
			if x >= btn.x && x <= btn.x+ButtonWidth && y >= btn.y && y <= btn.y+ButtonHeight {
				btn.action()
			}
		}
	}

	if !g.running {
		return nil
	}

	g.tickCount++
	if g.tickCount%UpdateSpeed != 0 {
		return nil
	}

	newGrid := make([][]bool, Rows)
	for i := range newGrid {
		newGrid[i] = make([]bool, Cols)
	}

	// Conway's Rules
	for i := range g.grid {
		for j := range g.grid[i] {
			neighbors := g.countNeighbors(i, j)
			if g.grid[i][j] {
				newGrid[i][j] = neighbors == 2 || neighbors == 3
			} else {
				newGrid[i][j] = neighbors == 3
			}
		}
	}
	g.grid = newGrid
	g.generation++
	return nil
}

func (g *Game) countNeighbors(x, y int) int {
	count := 0
	dirs := []struct{ dx, dy int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, d := range dirs {
		nx, ny := (x+d.dx+Rows)%Rows, (y+d.dy+Cols)%Cols
		if g.grid[nx][ny] {
			count++
		}
	}
	return count
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j] {
				x, y := j*CellSize, i*CellSize
				ebitenutil.DrawRect(screen, float64(x), float64(y), CellSize-1, CellSize-1, color.White)
			}
		}
	}
	msg := fmt.Sprintf("Generation: %d", g.generation)
	ebitenutil.DebugPrint(screen, msg)

	for _, btn := range g.buttons {
		drawButton(screen, btn.label, btn.x, btn.y)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	ScreenWidth, ScreenHeight = outsideWidth, outsideHeight
	Cols, Rows = ScreenWidth/CellSize, ScreenHeight/CellSize
	return ScreenWidth, ScreenHeight
}

func (g *Game) ToggleRunning() {
	g.running = !g.running

	for i := range g.buttons {
		if g.buttons[i].label == "Start" || g.buttons[i].label == "Pause" {
			if g.running {
				g.buttons[i].label = "Pause"
			} else {
				g.buttons[i].label = "Start"
			}
		}
	}
}

func (g *Game) Reset() {
	g.grid = make([][]bool, Rows)
	for i := range g.grid {
		g.grid[i] = make([]bool, Cols)
	}
	g.generation = 0
}

func (g *Game) Randomize() {
	rand.Seed(time.Now().UnixNano())
	for i := range g.grid {
		for j := range g.grid[i] {
			g.grid[i][j] = rand.Float64() < 0.2
		}
	}
}

func drawButton(screen *ebiten.Image, text string, x, y int) {
	ebitenutil.DrawRect(screen, float64(x), float64(y), ButtonWidth, ButtonHeight, color.White)
	ebitenutil.DebugPrintAt(screen, text, x+10, y+10)
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
