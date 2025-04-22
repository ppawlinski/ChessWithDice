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

func (c *Coordinates) Valid() bool {
	return c.col >= 0 && c.col < maxDimensions && c.row >= 0 && c.row < maxDimensions
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
	fields       [maxDimensions][maxDimensions]MovableDrawable
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
				piece.Draw(screen, x, y)
			}
		}
	}
}

func (b *Board) Reset() {
	//!!!3PPA todo change direction (orientation)
	b.fields[0][0] = &Rook{NewPiece(Black)}
	b.fields[1][0] = &Knight{NewPiece(Black)}
	b.fields[2][0] = &Bishop{NewPiece(Black)}
	b.fields[3][0] = &King{NewPiece(Black)}
	b.fields[4][0] = &Queen{NewPiece(Black)}
	b.fields[5][0] = &Bishop{NewPiece(Black)}
	b.fields[6][0] = &Knight{NewPiece(Black)}
	b.fields[7][0] = &Rook{NewPiece(Black)}

	b.fields[0][7] = &Rook{NewPiece(White)}
	b.fields[1][7] = &Knight{NewPiece(White)}
	b.fields[2][7] = &Bishop{NewPiece(White)}
	b.fields[3][7] = &King{NewPiece(White)}
	b.fields[4][7] = &Queen{NewPiece(White)}
	b.fields[5][7] = &Bishop{NewPiece(White)}
	b.fields[6][7] = &Knight{NewPiece(White)}
	b.fields[7][7] = &Rook{NewPiece(White)}

	b.fields[0][1] = &Pawn{NewPiece(Black)}
	b.fields[1][1] = &Pawn{NewPiece(Black)}
	b.fields[2][1] = &Pawn{NewPiece(Black)}
	b.fields[3][1] = &Pawn{NewPiece(Black)}
	b.fields[4][1] = &Pawn{NewPiece(Black)}
	b.fields[5][1] = &Pawn{NewPiece(Black)}
	b.fields[6][1] = &Pawn{NewPiece(Black)}
	b.fields[7][1] = &Pawn{NewPiece(Black)}
	b.fields[0][6] = &Pawn{NewPiece(White)}
	b.fields[1][6] = &Pawn{NewPiece(White)}
	b.fields[2][6] = &Pawn{NewPiece(White)}
	b.fields[3][6] = &Pawn{NewPiece(White)}
	b.fields[4][6] = &Pawn{NewPiece(White)}
	b.fields[5][6] = &Pawn{NewPiece(White)}
	b.fields[6][6] = &Pawn{NewPiece(White)}
	b.fields[7][6] = &Pawn{NewPiece(White)}

	b.kingPosition = [2]Coordinates{{3, 7}, {3, 0}}
}

func (b *Board) HitCheck(state *GameState) Coordinates {
	resultCol := -1
	resultRow := -1
	x, y := ebiten.CursorPosition()
	col := x / (config.TileSize + tileOffset)
	row := y / (config.TileSize + tileOffset)
	if x > 0 && y > 0 && col >= 0 && col < maxDimensions && row >= 0 && row < maxDimensions {
		/* if state.selectedPiece.Undefined(){
			if b.fields[col][row] != nil && b.fields[col][row].Piece().color == state.colorToMove {
				resultCol = col
				resultRow = row
				piece := b.fields[col][row].Piece()
				piece.SetDragOffset(x, y)
				fmt.Println("Selected piece: ", piece)
				state.possibleMoves = b.fields[col][row].GetPossibleMoves(b, Normal, Coordinates{resultCol, resultRow})
			}
		} */
		resultCol = col
		resultRow = row
	}

	return Coordinates{resultCol, resultRow}
}

func (b *Board) SelectPiece(s *GameState) {
	selectedPiece := Coordinates{-1, -1}
	field := b.HitCheck(s)
	fmt.Println("selected", field)
	fmt.Println(s.colorToMove)
	if b.GetPiece(field) != nil && b.GetPiece(field).Piece().color == s.colorToMove {
		selectedPiece = field
	}
	s.selectedPiece = selectedPiece
}

func (b *Board) GetPossibleMoves(c Coordinates) []Coordinates {
	//piece := b.fields[c.col][c.row].Piece()

	return b.fields[c.col][c.row].GetPossibleMoves(b, Normal, Coordinates{c.col, c.row})

}

func (b *Board) MoveSelected(state *GameState, selectedMove Coordinates) {
	dropped := b.fields[state.selectedPiece.col][state.selectedPiece.row].Piece()
	dropped.SetDragOffset(0, 0)

	moves := state.possibleMoves
	for _, move := range moves {
		if move.col == selectedMove.col && move.row == selectedMove.row {
			state.possibleMoves = nil
			dropped.firstMove = false
			state.colorToMove = (state.colorToMove + 1) % 2
			b.fields[move.col][move.row] = b.fields[state.selectedPiece.col][state.selectedPiece.row]
			if b.fields[move.col][move.row].Type() == KingType {
				b.kingPosition[b.fields[move.col][move.row].Piece().color] = move
			}
			b.fields[state.selectedPiece.col][state.selectedPiece.row] = nil
			state.selectedPiece.Reset()
			break
		}
	}
}

func (b *Board) GetPiece(position Coordinates) MovableDrawable {
	return b.fields[position.col][position.row]
}

func (b *Board) GetColor(position Coordinates) Color {
	var color Color

	if position.Valid() && b.fields[position.col][position.row] != nil {
		color = b.fields[position.col][position.row].Piece().color
	}

	return color
}
