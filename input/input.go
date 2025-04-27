package chessInput

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type InputEvent int8

const (
	Idle InputEvent = iota
	Click
	Release
	Hold
	InputEventLength int = iota
)

type Button int8

const (
	LMB Button = iota
	RMB
	ButtonLength int = iota
)

type Input struct {
	lmbPressed bool
	rmbPressed bool
}

func NewInput() *Input {
	input := &Input{lmbPressed: false, rmbPressed: false}
	return input
}

func (i *Input) GetButtonEvent(b Button) InputEvent {
	event := Idle
	previousState := i.getPreviousButtonState(b)
	currentState := i.getCurrentButtonState(b)
	if previousState {
		if currentState {
			event = Hold
		} else {
			event = Release
		}
	} else {
		if currentState {
			event = Click
		}
	}

	i.setPreviousButtonState(b, currentState)

	return event
}

func (i *Input) getPreviousButtonState(b Button) bool {
	var state bool
	switch b {
	case LMB:
		state = i.lmbPressed
	case RMB:
		state = i.rmbPressed
	}

	return state
}

func (i *Input) getCurrentButtonState(b Button) bool {
	var state bool
	switch b {
	case LMB:
		state = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	case RMB:
		state = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	}

	return state
}

func (i *Input) setPreviousButtonState(b Button, state bool) {
	switch b {
	case LMB:
		i.lmbPressed = state
	case RMB:
		i.rmbPressed = state
	}
}
