package main

type GameState struct {
	colorToMove   Color
	selectedPiece Coordinates
	dragging      bool
	possibleMoves []Coordinates
	//todo currentBudget  int
	leftoverBudget []bool
}

func NewGameState() *GameState {
	return &GameState{
		colorToMove:    White,
		selectedPiece:  Coordinates{-1, -1},
		leftoverBudget: []bool{false, false},
	}
}
