package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	chessInput "github.com/ppawlinski/ChessWithDice/input"
	"golang.org/x/image/font/gofont/goregular"
)

type Chess struct {
	board *Board
	state *GameState
	input *chessInput.Input
	ui    *ebitenui.UI
}

func NewChess() *Chess {
	chess := Chess{
		board: NewBoard(),
		state: NewGameState(),
		input: chessInput.NewInput(),
	}
	chess.board.Reset()
	return &chess
}

func (c *Chess) CreateUI() *ebitenui.UI {
	face, _ := loadFont(15)
	endTurnButtonImage, _ := loadButtonImage()
	var rollButton *widget.Button
	var endTurnButton *widget.Button
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(
					widget.NewInsetsSimple(10),
				),
			),
		),
	)
	endTurnButton = widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionEnd,
					VerticalPosition:   widget.AnchorLayoutPositionEnd,
				},
			),
		),
		widget.ButtonOpts.Image(endTurnButtonImage),
		widget.ButtonOpts.Text(
			"End turn",
			face,
			&widget.ButtonTextColor{
				Idle: color.RGBA{0, 0, 0, 255},
			},
		),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			rollButton.GetWidget().Visibility = widget.Visibility_Show
			endTurnButton.GetWidget().Visibility = widget.Visibility_Hide
		}),
	)

	rollButton = widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionEnd,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
		widget.ButtonOpts.Image(endTurnButtonImage),
		widget.ButtonOpts.Text(
			"Roll",
			face,
			&widget.ButtonTextColor{
				Idle: color.RGBA{0, 0, 0, 255},
			},
		),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    30,
			Bottom: 30,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			rollButton.GetWidget().Visibility = widget.Visibility_Hide
			endTurnButton.GetWidget().Visibility = widget.Visibility_Show
		}),
		//!!!3PPA visibility widget.WidgetOpts.
	)

	rootContainer.AddChild(endTurnButton)
	rootContainer.AddChild(rollButton)
	return &ebitenui.UI{
		Container: rootContainer,
	}
}

func (c *Chess) Update() error {
	c.ui.Update()
	lmbEvent := c.input.GetButtonEvent(chessInput.LMB)
	if lmbEvent == chessInput.Click {
		if c.state.selectedPiece.Undefined() {
			c.board.SelectPiece(c.state)
			if !c.state.selectedPiece.Undefined() {
				c.state.dragging = true
				c.state.possibleMoves = c.board.GetPossibleMoves(c.state.selectedPiece)
			}
		} else {
			selectedMove := c.board.HitCheck(c.state)
			if c.board.GetColor(selectedMove) == c.state.colorToMove {
				c.board.SelectPiece(c.state)
				c.state.dragging = true
				c.state.possibleMoves = c.board.GetPossibleMoves(c.state.selectedPiece)
			} else {
				c.board.MoveSelected(c.state, selectedMove)
			}
		}
	} else if lmbEvent == chessInput.Hold && c.state.dragging {
		c.board.GetPiece(c.state.selectedPiece).Piece().SetDragOffset(ebiten.CursorPosition())
	} else if lmbEvent == chessInput.Release {
		if c.state.dragging {
			selectedMove := c.board.HitCheck(c.state)
			c.board.MoveSelected(c.state, selectedMove)
		}
		c.state.dragging = false
	}
	return nil
}

func (c *Chess) Draw(screen *ebiten.Image) {
	c.board.Draw(screen, c.state)
	c.ui.Draw(screen)
}

func (g *Chess) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 806, 606 //!!!3PPA TODO embed it in some UI element
}
func loadFont(size float64) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Error loading font: %w", err)
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}

func loadButtonImage() (*widget.ButtonImage, error) {

	idle := image.NewBorderedNineSliceColor(color.NRGBA{90, 230, 76, 255}, color.NRGBA{90, 90, 90, 255}, 3)

	hover := image.NewBorderedNineSliceColor(color.NRGBA{70, 170, 76, 255}, color.NRGBA{70, 70, 70, 255}, 3)

	pressed := image.NewAdvancedNineSliceColor(color.NRGBA{70, 120, 96, 255}, image.NewBorder(3, 2, 2, 2, color.NRGBA{70, 70, 70, 255}))

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
