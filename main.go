package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

var chess = NewChess()
var draggedPiece Coordinates = Coordinates{-1, -1}
var possibleMoves []Coordinates

type Game struct{}
type Draggable interface {
	SetDragOffset(x, y int)
}

func (g *Game) Update(screen *ebiten.Image) error {
	return chess.Update(screen)
}

func (g *Game) Draw(screen *ebiten.Image) {
	chess.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 606, 606
}

func main() {
	ebiten.SetWindowSize(606, 606)
	ebiten.SetWindowTitle("Chess 2.0")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
