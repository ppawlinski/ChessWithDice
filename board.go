package main

import (
	"fmt"

	"github.com/ppawlinski/ChessWithDice/assets"
	"github.com/ppawlinski/ChessWithDice/config"

	"github.com/hajimehoshi/ebiten"
)

const (
	maxDimensions = 8
	tileOffset    = 1
)

type Draggable interface {
	SetDragOffset(x, y int)
}

type Coordinates struct {
	col int
	row int
}

func (c *Coordinates) Undefined() bool {
	return c.col == -1 && c.row == -1
}

func (c *Coordinates) Reset() {
	c.col = -1
	c.row = -1
}

type Movable interface {
	Move(Coordinates)
	GetPossibleMoves(*Board, Direction, Coordinates) []Coordinates
	Type() PieceType
	Piece() *Piece
}

type Drawable interface {
	Draw(*ebiten.Image, int, int)
}

type MovableDrawable interface {
	Movable
	Drawable
}

type Board struct {
	fields       [maxDimensions][maxDimensions]*Field
	kingPosition [ColorLength]Coordinates
	direction    Direction
}

func NewBoard() *Board {
	return &Board{
		kingPosition: [2]Coordinates{{1, 1}, {1, 1}},
		direction:    Normal,
	}
}

func (b *Board) Draw(screen *ebiten.Image, state GameState) {
	var rect *ebiten.Image
	for row := 0; row < maxDimensions; row++ {
		for col := 0; col < maxDimensions; col++ {
			if (row%2 == 1 && col%2 == 1) || (row%2 == 0 && col%2 == 0) {
				rect = assets.Images.DarkSquare
			} else {
				rect = assets.Images.LightSquare
			}
			for i := 0; i < len(state.possibleMoves); i++ {
				currentCoordinates := Coordinates{col, row}
				if state.possibleMoves[i] == currentCoordinates {
					rect = assets.Images.HighlightedSquare
				}
			}
			x := col * (config.TileSize + tileOffset)
			y := row * (config.TileSize + tileOffset)
			drawOptions := ebiten.DrawImageOptions{}
			drawOptions.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(rect, &drawOptions)
		}
	}

	for row := 0; row < maxDimensions; row++ {
		for col := 0; col < maxDimensions; col++ {
			x := col * (config.TileSize + tileOffset)
			y := row * (config.TileSize + tileOffset)
			piece := b.fields[col][row]
			if piece != nil {
				piece.piece.Draw(screen, x, y)
			}
		}
	}
}

func (b *Board) Reset() {
	//!!!3PPA todo change direction (orientation)
	b.fields[0][0] = &Field{piece: &Rook{NewPiece(Black)}}
	b.fields[1][0] = &Field{piece: &Knight{NewPiece(Black)}}
	b.fields[2][0] = &Field{piece: &Bishop{NewPiece(Black)}}
	b.fields[3][0] = &Field{piece: &King{NewPiece(Black)}}
	b.fields[4][0] = &Field{piece: &Queen{NewPiece(Black)}}
	b.fields[5][0] = &Field{piece: &Bishop{NewPiece(Black)}}
	b.fields[6][0] = &Field{piece: &Knight{NewPiece(Black)}}
	b.fields[7][0] = &Field{piece: &Rook{NewPiece(Black)}}

	b.fields[0][7] = &Field{piece: &Rook{NewPiece(White)}}
	b.fields[1][7] = &Field{piece: &Knight{NewPiece(White)}}
	b.fields[2][7] = &Field{piece: &Bishop{NewPiece(White)}}
	b.fields[3][7] = &Field{piece: &King{NewPiece(White)}}
	b.fields[4][7] = &Field{piece: &Queen{NewPiece(White)}}
	b.fields[5][7] = &Field{piece: &Bishop{NewPiece(White)}}
	b.fields[6][7] = &Field{piece: &Knight{NewPiece(White)}}
	b.fields[7][7] = &Field{piece: &Rook{NewPiece(White)}}

	b.fields[0][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[1][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[2][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[3][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[4][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[5][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[6][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[7][1] = &Field{piece: &Pawn{NewPiece(Black)}}
	b.fields[0][6] = &Field{piece: &Pawn{NewPiece(White)}}
	b.fields[1][6] = &Field{piece: &Pawn{NewPiece(White)}}
	b.fields[2][6] = &Field{piece: &Pawn{NewPiece(White)}}
	b.fields[3][6] = &Field{piece: &Pawn{NewPiece(White)}}
	b.fields[4][6] = &Field{piece: &Pawn{NewPiece(White)}}
	b.fields[5][6] = &Field{piece: &Pawn{NewPiece(White)}}
	b.fields[6][6] = &Field{piece: &Pawn{NewPiece(White)}}
	b.fields[7][6] = &Field{piece: &Pawn{NewPiece(White)}}

	b.kingPosition = [2]Coordinates{{3, 7}, {3, 0}}
}

func (b *Board) HitCheck(state *GameState) (Coordinates, []Coordinates) {
	resultCol := -1
	resultRow := -1
	x, y := ebiten.CursorPosition()
	col := x / (config.TileSize + tileOffset)
	row := y / (config.TileSize + tileOffset)
	if x > 0 && y > 0 && col >= 0 && col < maxDimensions && row >= 0 && row < maxDimensions {
		if b.fields[col][row] != nil && b.fields[col][row].piece.Piece().color == state.colorToMove {
			resultCol = col
			resultRow = row
			piece := b.fields[col][row].piece.Piece()
			piece.SetDragOffset(x, y)
			fmt.Println("Selected piece: ", piece)
			state.possibleMoves = b.fields[col][row].piece.GetPossibleMoves(b, Normal, Coordinates{resultCol, resultRow})
		}
	}

	return Coordinates{resultCol, resultRow}, state.possibleMoves
}

func (b *Board) DropCheck(state *GameState) {
	draggedPiece := state.selectedPiece
	x, y := ebiten.CursorPosition()
	col := x / (config.TileSize + tileOffset)
	row := y / (config.TileSize + tileOffset)

	dropped := b.fields[draggedPiece.col][draggedPiece.row].piece.Piece()
	dropped.SetDragOffset(0, 0)
	state.possibleMoves = nil

	moves := b.fields[draggedPiece.col][draggedPiece.row].piece.GetPossibleMoves(b, Normal, draggedPiece)
	for _, move := range moves {
		if move.col == col && move.row == row {
			dropped.firstMove = false
			state.colorToMove = (state.colorToMove + 1) % 2
			b.fields[col][row] = b.fields[draggedPiece.col][draggedPiece.row]
			if b.fields[col][row].piece.Type() == KingType {
				b.kingPosition[b.fields[col][row].piece.Piece().color] = move
			}
			b.fields[draggedPiece.col][draggedPiece.row] = nil
			break
		}
	}
}

func (b *Board) GetPiece(position Coordinates) Draggable {
	return b.fields[position.col][position.row].piece.Piece()
}
