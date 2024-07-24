package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Color int8

const (
	White Color = iota
	Black
	ColorLength int = iota
)

type PieceType int8

type Offset struct {
	x, y int
}

const (
	PawnType PieceType = iota
	RookType
	KnightType
	BishopType
	QueenType
	KingType
)

type Piece struct {
	color      Color
	dragOffset Offset
	firstMove  bool
}

func NewPiece(color Color) *Piece {
	return &Piece{
		color:      color,
		dragOffset: Offset{},
		firstMove:  true,
	}
}

func (p *Piece) Draw(screen *ebiten.Image, x int, y int, path string) {
	filter := ebiten.FilterDefault
	options := ebiten.DrawImageOptions{}
	image, _, _ := ebitenutil.NewImageFromFile(path, filter)
	imageX, imageY := image.Size()
	offsetX := (tileSize - imageX) / 2
	offsetY := (tileSize - imageY) / 2
	if !(p.dragOffset.x == 0 && p.dragOffset.y == 0) {
		x = p.dragOffset.x
		y = p.dragOffset.y
		offsetX -= tileSize / 2
		offsetY -= tileSize / 2
	}
	options.GeoM.Translate(float64(x+offsetX), float64(y+offsetY))
	screen.DrawImage(image, &options)

}

func (p *Piece) SetDragOffset(x, y int) {
	p.dragOffset = Offset{x, y}
}

type Pawn struct {
	piece *Piece
}

func (p *Pawn) Move(Coordinates) {
}

func (p *Pawn) Draw(screen *ebiten.Image, x int, y int) {
	path := "images\\pawn.png"
	if p.piece.color == White {
		path = "images\\pawnWhite.png"
	}

	p.piece.Draw(screen, x, y, path)
}

func (p *Pawn) Type() PieceType {
	return PawnType
}

func (p *Pawn) Piece() *Piece {
	return p.piece
}

type Rook struct {
	piece *Piece
}

func (r *Rook) Move(Coordinates) {

}

func (r *Rook) Draw(screen *ebiten.Image, x int, y int) {
	path := "images\\rook.png"
	if r.piece.color == White {
		path = "images\\rookWhite.png"
	}
	r.piece.Draw(screen, x, y, path)
}

func (r *Rook) Type() PieceType {
	return RookType
}

func (r *Rook) Piece() *Piece {
	return r.piece
}

type Knight struct {
	piece *Piece
}

func (k *Knight) Move(Coordinates) {

}

func (k *Knight) Draw(screen *ebiten.Image, x int, y int) {
	path := "images\\knight.png"

	if k.piece.color == White {
		path = "images\\knightWhite.png"
	}

	k.piece.Draw(screen, x, y, path)

}

func (k *Knight) Type() PieceType {
	return KnightType
}

func (k *Knight) Piece() *Piece {
	return k.piece
}

type Bishop struct {
	piece *Piece
}

func (b *Bishop) Move(Coordinates) {

}

func (b *Bishop) Draw(screen *ebiten.Image, x int, y int) {
	path := "images\\bishop.png"
	if b.piece.color == White {
		path = "images\\bishopWhite.png"
	}
	b.piece.Draw(screen, x, y, path)
}

func (b *Bishop) Type() PieceType {
	return BishopType
}

func (b *Bishop) Piece() *Piece {
	return b.piece
}

type Queen struct {
	piece *Piece
}

func (q *Queen) Move(Coordinates) {

}

func (q *Queen) Draw(screen *ebiten.Image, x int, y int) {
	path := "images\\queen.png"
	if q.piece.color == White {
		path = "images\\queenWhite.png"
	}
	q.piece.Draw(screen, x, y, path)
}

func (q *Queen) Type() PieceType {
	return QueenType
}

func (q *Queen) Piece() *Piece {
	return q.piece
}

type King struct {
	piece *Piece
}

func (k *King) Move(Coordinates) {

}

func (k *King) Draw(screen *ebiten.Image, x int, y int) {
	path := "images\\king.png"
	if k.piece.color == White {
		path = "images\\kingWhite.png"
	}
	k.piece.Draw(screen, x, y, path)
}

func (k *King) Type() PieceType {
	return KingType
}

func (k *King) Piece() *Piece {
	return k.piece
}
