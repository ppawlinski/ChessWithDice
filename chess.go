package main

import "github.com/hajimehoshi/ebiten"

type Chess struct {
	board Board
	state GameState
}

func NewChess() *Chess {
	chess := Chess{board: *NewBoard(), state: *NewGameState()}
	chess.board.Reset()

	return &chess
}

func (c *Chess) Update(screen *ebiten.Image) error {
	//!!!3PPA todo add some logic to determine if it's initial mouseclick - prevent catching a piece mid-drag
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if c.state.dragging {
			c.board.GetPiece(c.state.selectedPiece).SetDragOffset(ebiten.CursorPosition())
		} else {
			c.state.selectedPiece, c.state.possibleMoves = c.board.HitCheck(&c.state)
			if !c.state.selectedPiece.Undefined() {
				c.state.dragging = true
			}
		}
	} else if c.state.dragging {
		c.board.DropCheck(&c.state)
		c.state.selectedPiece.Reset()
		c.state.dragging = false
	}
	return nil
}

func (c *Chess) Draw(screen *ebiten.Image) {
	c.board.Draw(screen, c.state)
}

func (g *Chess) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 606, 606
}
