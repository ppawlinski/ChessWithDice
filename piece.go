package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ppawlinski/ChessWithDice/assets"
	"github.com/ppawlinski/ChessWithDice/config"
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

func (p *Piece) Draw(screen *ebiten.Image, x int, y int, image *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	imageX, imageY := image.Size()
	offsetX := (config.TileSize - imageX) / 2
	offsetY := (config.TileSize - imageY) / 2
	if !(p.dragOffset.x == 0 && p.dragOffset.y == 0) {
		x = p.dragOffset.x
		y = p.dragOffset.y
		offsetX -= config.TileSize / 2
		offsetY -= config.TileSize / 2
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
	image := assets.Images.BlackPawn
	if p.piece.color == White {
		image = assets.Images.WhitePawn
	}

	p.piece.Draw(screen, x, y, image)
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
	image := assets.Images.BlackRook
	if r.piece.color == White {
		image = assets.Images.WhiteRook
	}
	r.piece.Draw(screen, x, y, image)
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
	image := assets.Images.BlackKnight

	if k.piece.color == White {
		image = assets.Images.WhiteKnight
	}

	k.piece.Draw(screen, x, y, image)

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
	image := assets.Images.BlackBishop
	if b.piece.color == White {
		image = assets.Images.WhiteBishop
	}
	b.piece.Draw(screen, x, y, image)
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
	image := assets.Images.BlackQueen
	if q.piece.color == White {
		image = assets.Images.WhiteQueen
	}
	q.piece.Draw(screen, x, y, image)
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
	image := assets.Images.BlackKing
	if k.piece.color == White {
		image = assets.Images.WhiteKing
	}
	k.piece.Draw(screen, x, y, image)
}

func (k *King) Type() PieceType {
	return KingType
}

func (k *King) Piece() *Piece {
	return k.piece
}
