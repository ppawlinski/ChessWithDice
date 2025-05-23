package main

type GameState struct {
	turnStarted    bool
	colorToMove    Color
	selectedPiece  Coordinates
	dragging       bool
	possibleMoves  []Coordinates
	currentBudget  int
	leftoverBudget []bool
	enPassant      Coordinates
}

func NewGameState() *GameState {
	return &GameState{
		turnStarted:    false,
		colorToMove:    White,
		selectedPiece:  Coordinates{-1, -1},
		leftoverBudget: []bool{false, false},
		currentBudget:  0,
		enPassant:      Coordinates{-1, -1},
	}
}
