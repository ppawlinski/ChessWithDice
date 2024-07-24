package main

type Direction int8

const (
	Normal Direction = iota
	Reversed
)

func (p *Pawn) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	moveDirection := 1
	if direction == Reversed {
		moveDirection *= -1
	}
	if p.piece.color == White {
		moveDirection *= -1
	}

	if current.row+moveDirection < maxDimensions && current.row+moveDirection >= 0 {
		if b.fields[current.col][current.row+moveDirection] == nil {
			moves = append(moves, Coordinates{current.col, current.row + moveDirection})
			if p.piece.firstMove {
				if b.fields[current.col][current.row+(moveDirection*2)] == nil {
					moves = append(moves, Coordinates{current.col, current.row + (moveDirection * 2)})

				}
			}
		}
		if current.col-1 >= 0 && b.fields[current.col-1][current.row+moveDirection] != nil {
			if b.fields[current.col-1][current.row+moveDirection].piece.Piece().color != p.piece.color {
				moves = append(moves, Coordinates{current.col - 1, current.row + moveDirection})
			}
		}
		if current.col+1 < maxDimensions && b.fields[current.col+1][current.row+moveDirection] != nil {
			if b.fields[current.col+1][current.row+moveDirection].piece.Piece().color != p.piece.color {
				moves = append(moves, Coordinates{current.col + 1, current.row + moveDirection})
			}
		}
	}

	return FilterIllegalMoves(current, moves, b, p.piece.color)
}

func (r *Rook) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].piece.Piece().color
	moves = getRookMoves(current, b, currentColor, moves)
	return FilterIllegalMoves(current, moves, b, currentColor)
}

func getRookMoves(current Coordinates, b *Board, currentColor Color, moves []Coordinates) []Coordinates {
	for i := 1; current.col+i < maxDimensions; i++ {
		potentialField := b.fields[current.col+i][current.row]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col + i, current.row})
		}
		if potentialField != nil {
			break
		}
	}
	for i := 1; current.col-i >= 0; i++ {
		potentialField := b.fields[current.col-i][current.row]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col - i, current.row})
		}
		if potentialField != nil {
			break
		}
	}
	for i := 1; current.row+i < maxDimensions; i++ {
		potentialField := b.fields[current.col][current.row+i]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col, current.row + i})
		}
		if potentialField != nil {
			break
		}
	}
	for i := 1; current.row-i >= 0; i++ {
		potentialField := b.fields[current.col][current.row-i]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col, current.row - i})
		}
		if potentialField != nil {
			break
		}
	}

	return FilterIllegalMoves(current, moves, b, currentColor)
}

func (k *Knight) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].piece.Piece().color

	if current.col+2 < maxDimensions {
		if current.row+1 < maxDimensions {
			if b.fields[current.col+2][current.row+1] == nil ||
				b.fields[current.col+2][current.row+1].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 2, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			if b.fields[current.col+2][current.row-1] == nil ||
				b.fields[current.col+2][current.row-1].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 2, current.row - 1})
			}
		}
	}
	if current.col-2 >= 0 {
		if current.row+1 < maxDimensions {
			if b.fields[current.col-2][current.row+1] == nil ||
				b.fields[current.col-2][current.row+1].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 2, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			if b.fields[current.col-2][current.row-1] == nil ||
				b.fields[current.col-2][current.row-1].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 2, current.row - 1})
			}
		}
	}
	if current.col+1 < maxDimensions {
		if current.row+2 < maxDimensions {
			if b.fields[current.col+1][current.row+2] == nil ||
				b.fields[current.col+1][current.row+2].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row + 2})
			}
		}
		if current.row-2 >= 0 {
			if b.fields[current.col+1][current.row-2] == nil ||
				b.fields[current.col+1][current.row-2].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row - 2})
			}
		}
	}
	if current.col-1 >= 0 {
		if current.row+2 < maxDimensions {
			if b.fields[current.col-1][current.row+2] == nil ||
				b.fields[current.col-1][current.row+2].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row + 2})
			}
		}
		if current.row-2 >= 0 {
			if b.fields[current.col-1][current.row-2] == nil ||
				b.fields[current.col-1][current.row-2].piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row - 2})
			}
		}
	}

	return FilterIllegalMoves(current, moves, b, currentColor)
}

func (b *Bishop) GetPossibleMoves(board *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := board.fields[current.col][current.row].piece.Piece().color
	moves = getBishopMoves(current, board, currentColor, moves)

	return FilterIllegalMoves(current, moves, board, currentColor)
}

func getBishopMoves(current Coordinates, board *Board, currentColor Color, moves []Coordinates) []Coordinates {
	for i := 1; current.col+i < maxDimensions && current.row+i < maxDimensions; i++ {
		potentialField := board.fields[current.col+i][current.row+i]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col + i, current.row + i})
		}
		if potentialField != nil {
			break
		}
	}
	for i := 1; current.col-i >= 0 && current.row-i >= 0; i++ {
		potentialField := board.fields[current.col-i][current.row-i]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col - i, current.row - i})
		}
		if potentialField != nil {
			break
		}
	}
	for i := 1; current.col+i < maxDimensions && current.row-i >= 0; i++ {
		potentialField := board.fields[current.col+i][current.row-i]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col + i, current.row - i})
		}
		if potentialField != nil {
			break
		}
	}
	for i := 1; current.col-i >= 0 && current.row+i < maxDimensions; i++ {
		potentialField := board.fields[current.col-i][current.row+i]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col - i, current.row + i})
		}
		if potentialField != nil {
			break
		}
	}

	return FilterIllegalMoves(current, moves, board, currentColor)
}

func (q *Queen) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].piece.Piece().color
	moves = getBishopMoves(current, b, currentColor, moves)
	moves = getRookMoves(current, b, currentColor, moves)

	return FilterIllegalMoves(current, moves, b, currentColor)
}

func (k *King) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].piece.Piece().color
	if current.col+1 < maxDimensions {
		potentialField := b.fields[current.col+1][current.row]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col + 1, current.row})
		}
		if current.row+1 < maxDimensions {
			potentialField := b.fields[current.col+1][current.row+1]
			if potentialField == nil || potentialField.piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			potentialField := b.fields[current.col+1][current.row-1]
			if potentialField == nil || potentialField.piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row - 1})
			}
		}
	}
	if current.col-1 >= 0 {
		potentialField := b.fields[current.col-1][current.row]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col - 1, current.row})
		}
		if current.row+1 < maxDimensions {
			potentialField := b.fields[current.col-1][current.row+1]
			if potentialField == nil || potentialField.piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			potentialField := b.fields[current.col-1][current.row-1]
			if potentialField == nil || potentialField.piece.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row - 1})
			}
		}
	}
	if current.row+1 < maxDimensions {
		potentialField := b.fields[current.col][current.row+1]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col, current.row + 1})
		}
	}
	if current.row-1 >= 0 {
		potentialField := b.fields[current.col][current.row-1]
		if potentialField == nil || potentialField.piece.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col, current.row - 1})
		}
	}
	return moves
}

func FilterIllegalMoves(position Coordinates, moves []Coordinates, board *Board, color Color) []Coordinates {
	var legalMoves []Coordinates
	for _, move := range moves {
		if IsLegalMove(position, move, board, color) {
			legalMoves = append(legalMoves, move)
		}
	}
	return legalMoves
}

func IsLegalMove(position Coordinates, move Coordinates, board *Board, color Color) bool {
	kingPosition := board.kingPosition[color]
	var direction int
	legal := true

	if kingPosition.col == position.col && kingPosition.col != move.col {
		if kingPosition.row-position.row > 0 {
			direction = -1
		} else {
			direction = 1
		}

		for row := position.row + direction; row < maxDimensions || row >= 0; row += direction {
			field := board.fields[kingPosition.col][row]
			if field != nil {
				if field.piece.Piece().color != color && (field.piece.Type() == RookType || field.piece.Type() == QueenType) {
					legal = false
				}
				break
			}
		}
	}

	//if kingPosition.row == position.row && kingPosition.row != move.row {
	//}

	return legal
}
