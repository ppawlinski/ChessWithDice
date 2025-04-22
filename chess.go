package main

import (
	"github.com/hajimehoshi/ebiten"
	input "github.com/ppawlinski/ChessWithDice/input"
)

type Chess struct {
	board Board
	state GameState
	input input.Input
}

func NewChess() *Chess {
	chess := Chess{board: *NewBoard(), state: *NewGameState()}
	chess.board.Reset()

	return &chess
}

func (c *Chess) Update(screen *ebiten.Image) error {
	lmbEvent := c.input.GetButtonEvent(input.LMB)
	if lmbEvent == input.Click {
		if c.state.selectedPiece.Undefined() {
			c.board.SelectPiece(&c.state)
			if !c.state.selectedPiece.Undefined() {
				c.state.dragging = true
				c.state.possibleMoves = c.board.GetPossibleMoves(c.state.selectedPiece)
			}
		} else {
			selectedMove := c.board.HitCheck(&c.state)
			if c.board.GetColor(selectedMove) == c.state.colorToMove {
				c.board.SelectPiece(&c.state)
				c.state.dragging = true
				c.state.possibleMoves = c.board.GetPossibleMoves(c.state.selectedPiece)
			} else {
				c.board.MoveSelected(&c.state, selectedMove)
			}
		}
	} else if lmbEvent == input.Hold && c.state.dragging {
		c.board.GetPiece(c.state.selectedPiece).Piece().SetDragOffset(ebiten.CursorPosition())
	} else if lmbEvent == input.Release {
		if c.state.dragging {
			selectedMove := c.board.HitCheck(&c.state)
			c.board.MoveSelected(&c.state, selectedMove)
		}
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
