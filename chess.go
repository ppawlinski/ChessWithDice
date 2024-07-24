package main

import "github.com/hajimehoshi/ebiten"

type Chess struct {
	board Board
}

func NewChess() *Chess {
	chess := Chess{Board{kingPosition: [2]Coordinates{{1, 1}, {1, 1}}}}
	chess.board.Reset()

	return &chess
}

func (c *Chess) Update(screen *ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if draggedPiece.Undefined() {
			draggedPiece, possibleMoves = c.board.HitCheck()
		} else {
			c.board.GetPiece(draggedPiece).SetDragOffset(ebiten.CursorPosition())
		}
	} else if !draggedPiece.Undefined() {
		c.board.DropCheck(draggedPiece)
		draggedPiece.Reset()
	}
	return nil
}

func (c *Chess) Draw(screen *ebiten.Image) {
	c.board.Draw(screen)
}
