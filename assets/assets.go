package assets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/ppawlinski/ChessWithDice/config"
)

type Assets struct {
	DarkSquare        *ebiten.Image
	LightSquare       *ebiten.Image
	HighlightedSquare *ebiten.Image
	WhiteKing         *ebiten.Image
	BlackKing         *ebiten.Image
	WhiteQueen        *ebiten.Image
	BlackQueen        *ebiten.Image
	WhiteBishop       *ebiten.Image
	BlackBishop       *ebiten.Image
	WhiteKnight       *ebiten.Image
	BlackKnight       *ebiten.Image
	WhiteRook         *ebiten.Image
	BlackRook         *ebiten.Image
	WhitePawn         *ebiten.Image
	BlackPawn         *ebiten.Image
}

var Images Assets

func Init() {
	image, _ := ebiten.NewImage(config.TileSize, config.TileSize, ebiten.FilterDefault)
	image.Fill(color.RGBA{117, 59, 12, 255})
	Images.DarkSquare = image
	image, _ = ebiten.NewImage(config.TileSize, config.TileSize, ebiten.FilterDefault)
	image.Fill(color.RGBA{196, 151, 114, 255})
	Images.LightSquare = image
	image, _ = ebiten.NewImage(config.TileSize, config.TileSize, ebiten.FilterDefault)
	image.Fill(color.RGBA{255, 127, 80, 255})
	Images.HighlightedSquare = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\pawn.png", ebiten.FilterDefault)
	Images.BlackPawn = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\pawnWhite.png", ebiten.FilterDefault)
	Images.WhitePawn = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\king.png", ebiten.FilterDefault)
	Images.BlackKing = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\kingWhite.png", ebiten.FilterDefault)
	Images.WhiteKing = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\queen.png", ebiten.FilterDefault)
	Images.BlackQueen = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\queenWhite.png", ebiten.FilterDefault)
	Images.WhiteQueen = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\bishop.png", ebiten.FilterDefault)
	Images.BlackBishop = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\bishopWhite.png", ebiten.FilterDefault)
	Images.WhiteBishop = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\knight.png", ebiten.FilterDefault)
	Images.BlackKnight = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\knightWhite.png", ebiten.FilterDefault)
	Images.WhiteKnight = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\rook.png", ebiten.FilterDefault)
	Images.BlackRook = image
	image, _, _ = ebitenutil.NewImageFromFile("images\\rookWhite.png", ebiten.FilterDefault)
	Images.WhiteRook = image
}
