package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ppawlinski/ChessWithDice/assets"
)

func main() {
	assets.Init()
	ebiten.SetWindowSize(606, 606)
	ebiten.SetWindowTitle("Chess 2.0")
	if err := ebiten.RunGame(NewChess()); err != nil {
		log.Fatal(err)
	}
}
