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
		moves = getPawnMoves(b, current, moveDirection, moves, p)
		moves = getPawnTakes(current, b, moveDirection, p, moves)
		moves = getEnPassant(b, current, moveDirection, moves, p)
	}

	return FilterIllegalMoves(current, moves, b, p.piece.color)
}

func getEnPassant(b *Board, current Coordinates, moveDirection int, moves []Coordinates, p *Pawn) []Coordinates {
	if current.col-1 >= 0 {
		if other, ok := b.fields[current.col-1][current.row].(EnPassantable); ok {
			if other.EnPassantable() {
				moves = append(moves, Coordinates{current.col - 1, current.row + moveDirection})
			}
		}

	}

	if current.col+1 < maxDimensions {
		if other, ok := b.fields[current.col+1][current.row].(EnPassantable); ok {
			if other.EnPassantable() {
				moves = append(moves, Coordinates{current.col + 1, current.row + moveDirection})
			}
		}

	}

	return moves
}

func getPawnTakes(current Coordinates, b *Board, moveDirection int, p *Pawn, moves []Coordinates) []Coordinates {
	if current.col-1 >= 0 && b.fields[current.col-1][current.row+moveDirection] != nil {
		if b.fields[current.col-1][current.row+moveDirection].Piece().color != p.piece.color {
			moves = append(moves, Coordinates{current.col - 1, current.row + moveDirection})
		}
	}
	if current.col+1 < maxDimensions && b.fields[current.col+1][current.row+moveDirection] != nil {
		if b.fields[current.col+1][current.row+moveDirection].Piece().color != p.piece.color {
			moves = append(moves, Coordinates{current.col + 1, current.row + moveDirection})
		}
	}
	return moves
}

func getPawnMoves(b *Board, current Coordinates, moveDirection int, moves []Coordinates, p *Pawn) []Coordinates {
	if b.fields[current.col][current.row+moveDirection] == nil {
		moves = append(moves, Coordinates{current.col, current.row + moveDirection})
		if p.piece.firstMove {
			if b.fields[current.col][current.row+(moveDirection*2)] == nil {
				moves = append(moves, Coordinates{current.col, current.row + (moveDirection * 2)})

			}
		}
	}
	return moves
}

func (r *Rook) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].Piece().color
	moves = getRookMoves(current, b, currentColor, moves)
	return FilterIllegalMoves(current, moves, b, currentColor)
}

func getRookMoves(current Coordinates, b *Board, currentColor Color, moves []Coordinates) []Coordinates {
	moves = append(moves, GetMovesLeft(current, currentColor, b)...)
	moves = append(moves, GetMovesRight(current, currentColor, b)...)
	moves = append(moves, GetMovesDown(current, currentColor, b)...)
	moves = append(moves, GetMovesUp(current, currentColor, b)...)

	return moves
}

func (k *Knight) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].Piece().color

	if current.col+2 < maxDimensions {
		if current.row+1 < maxDimensions {
			if b.fields[current.col+2][current.row+1] == nil ||
				b.fields[current.col+2][current.row+1].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 2, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			if b.fields[current.col+2][current.row-1] == nil ||
				b.fields[current.col+2][current.row-1].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 2, current.row - 1})
			}
		}
	}
	if current.col-2 >= 0 {
		if current.row+1 < maxDimensions {
			if b.fields[current.col-2][current.row+1] == nil ||
				b.fields[current.col-2][current.row+1].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 2, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			if b.fields[current.col-2][current.row-1] == nil ||
				b.fields[current.col-2][current.row-1].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 2, current.row - 1})
			}
		}
	}
	if current.col+1 < maxDimensions {
		if current.row+2 < maxDimensions {
			if b.fields[current.col+1][current.row+2] == nil ||
				b.fields[current.col+1][current.row+2].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row + 2})
			}
		}
		if current.row-2 >= 0 {
			if b.fields[current.col+1][current.row-2] == nil ||
				b.fields[current.col+1][current.row-2].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row - 2})
			}
		}
	}
	if current.col-1 >= 0 {
		if current.row+2 < maxDimensions {
			if b.fields[current.col-1][current.row+2] == nil ||
				b.fields[current.col-1][current.row+2].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row + 2})
			}
		}
		if current.row-2 >= 0 {
			if b.fields[current.col-1][current.row-2] == nil ||
				b.fields[current.col-1][current.row-2].Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row - 2})
			}
		}
	}

	return FilterIllegalMoves(current, moves, b, currentColor)
}

func (b *Bishop) GetPossibleMoves(board *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := board.fields[current.col][current.row].Piece().color
	moves = getBishopMoves(current, board, currentColor, moves)

	return FilterIllegalMoves(current, moves, board, currentColor)
}

func getBishopMoves(current Coordinates, board *Board, currentColor Color, moves []Coordinates) []Coordinates {
	moves = getDiagonalMoves(1, 1, current, board, currentColor, moves)
	moves = getDiagonalMoves(-1, -1, current, board, currentColor, moves)
	moves = getDiagonalMoves(1, -1, current, board, currentColor, moves)
	moves = getDiagonalMoves(-1, 1, current, board, currentColor, moves)

	return moves
}

func getDiagonalMoves(colModifier, rowModifier int, current Coordinates, board *Board, currentColor Color, moves []Coordinates) []Coordinates {
	for i := 1; current.col+i*colModifier >= 0 && current.col+i*colModifier < maxDimensions && current.row+i*rowModifier >= 0 && current.row+i*rowModifier < maxDimensions; i++ {
		potentialField := board.fields[current.col+i*colModifier][current.row+i*rowModifier]
		if potentialField == nil || potentialField.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col + i*colModifier, current.row + i*rowModifier})
		}
		if potentialField != nil {
			break
		}
	}

	return moves
}

func (q *Queen) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].Piece().color
	moves = getBishopMoves(current, b, currentColor, moves)
	moves = getRookMoves(current, b, currentColor, moves)

	return FilterIllegalMoves(current, moves, b, currentColor)
}

func (k *King) GetPossibleMoves(b *Board, direction Direction, current Coordinates) []Coordinates {
	var moves []Coordinates
	currentColor := b.fields[current.col][current.row].Piece().color
	if current.col+1 < maxDimensions {
		potentialField := b.fields[current.col+1][current.row]
		if potentialField == nil || potentialField.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col + 1, current.row})
		}
		if current.row+1 < maxDimensions {
			potentialField := b.fields[current.col+1][current.row+1]
			if potentialField == nil || potentialField.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			potentialField := b.fields[current.col+1][current.row-1]
			if potentialField == nil || potentialField.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col + 1, current.row - 1})
			}
		}
	}
	if current.col-1 >= 0 {
		potentialField := b.fields[current.col-1][current.row]
		if potentialField == nil || potentialField.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col - 1, current.row})
		}
		if current.row+1 < maxDimensions {
			potentialField := b.fields[current.col-1][current.row+1]
			if potentialField == nil || potentialField.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row + 1})
			}
		}
		if current.row-1 >= 0 {
			potentialField := b.fields[current.col-1][current.row-1]
			if potentialField == nil || potentialField.Piece().color != currentColor {
				moves = append(moves, Coordinates{current.col - 1, current.row - 1})
			}
		}
	}
	if current.row+1 < maxDimensions {
		potentialField := b.fields[current.col][current.row+1]
		if potentialField == nil || potentialField.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col, current.row + 1})
		}
	}
	if current.row-1 >= 0 {
		potentialField := b.fields[current.col][current.row-1]
		if potentialField == nil || potentialField.Piece().color != currentColor {
			moves = append(moves, Coordinates{current.col, current.row - 1})
		}
	}

	return FilterIllegalMoves(current, moves, b, currentColor)
}

func FilterIllegalMoves(position Coordinates, moves []Coordinates, board *Board, color Color) []Coordinates {
	var legalMoves []Coordinates
	for _, move := range moves {
		if !KingInCheck(position, move, board, color) {
			legalMoves = append(legalMoves, move)
		}
	}
	return legalMoves
}

func KingInCheck(position Coordinates, move Coordinates, board *Board, color Color) bool {
	kingPosition := board.kingPosition[color]
	kingInCheck := false
	blockedField := Coordinates{-1, -1}

	if board.fields[position.col][position.row].Type() == KingType {
		kingPosition = move
	} else {
		blockedField = move
	}

	kingInCheck = AttackedByRookQueen(kingPosition, position, blockedField, board, color)
	kingInCheck = kingInCheck || AttackedByBishopQueen(kingPosition, position, blockedField, board, color)
	kingInCheck = kingInCheck || AttackedByPawn(kingPosition, position, blockedField, board, color)
	kingInCheck = kingInCheck || AttackedByKnight(kingPosition, position, blockedField, board, color)
	kingInCheck = kingInCheck || AttackedByKing(kingPosition, position, blockedField, board, color)

	return kingInCheck
}

func AttackedByKing(kingPosition, freeField, blockedField Coordinates, b *Board, color Color) bool {
	attacked := false
	if kingPosition.col+1 < maxDimensions {
		potentialField := b.fields[kingPosition.col+1][kingPosition.row]
		if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
			attacked = true
		}
		if kingPosition.row+1 < maxDimensions {
			potentialField := b.fields[kingPosition.col+1][kingPosition.row+1]
			if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
				attacked = true
			}
		}
		if kingPosition.row-1 >= 0 {
			potentialField := b.fields[kingPosition.col+1][kingPosition.row-1]
			if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
				attacked = true
			}
		}
	}
	if kingPosition.col-1 >= 0 {
		potentialField := b.fields[kingPosition.col-1][kingPosition.row]
		if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
			attacked = true
		}
		if kingPosition.row+1 < maxDimensions {
			potentialField := b.fields[kingPosition.col-1][kingPosition.row+1]
			if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
				attacked = true
			}
		}
		if kingPosition.row-1 >= 0 {
			potentialField := b.fields[kingPosition.col-1][kingPosition.row-1]
			if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
				attacked = true
			}
		}
	}
	if kingPosition.row+1 < maxDimensions {
		potentialField := b.fields[kingPosition.col][kingPosition.row+1]
		if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
			attacked = true
		}
	}
	if kingPosition.row-1 >= 0 {
		potentialField := b.fields[kingPosition.col][kingPosition.row-1]
		if potentialField != nil && potentialField.Piece().color != color && potentialField.Type() == KingType {
			attacked = true
		}
	}
	return attacked
}

func AttackedByKnight(kingPosition, freeField, blockedField Coordinates, board *Board, color Color) bool {
	attacked := false

	if kingPosition.col+2 < maxDimensions {
		if kingPosition.row+1 < maxDimensions {
			if board.fields[kingPosition.col+2][kingPosition.row+1] != nil &&
				board.fields[kingPosition.col+2][kingPosition.row+1].Piece().color != color &&
				board.fields[kingPosition.col+2][kingPosition.row+1].Type() == KnightType {
				attacked = true
			}
		}
		if kingPosition.row-1 >= 0 {
			if board.fields[kingPosition.col+2][kingPosition.row-1] != nil &&
				board.fields[kingPosition.col+2][kingPosition.row-1].Piece().color != color &&
				board.fields[kingPosition.col+2][kingPosition.row-1].Type() == KnightType {
				attacked = true

			}
		}
	}
	if kingPosition.col-2 >= 0 {
		if kingPosition.row+1 < maxDimensions {
			if board.fields[kingPosition.col-2][kingPosition.row+1] != nil &&
				board.fields[kingPosition.col-2][kingPosition.row+1].Piece().color != color &&
				board.fields[kingPosition.col-2][kingPosition.row+1].Type() == KnightType {
				attacked = true
			}
		}
		if kingPosition.row-1 >= 0 {
			if board.fields[kingPosition.col-2][kingPosition.row-1] != nil &&
				board.fields[kingPosition.col-2][kingPosition.row-1].Piece().color != color &&
				board.fields[kingPosition.col-2][kingPosition.row-1].Type() == KnightType {
				attacked = true
			}
		}
	}
	if kingPosition.col+1 < maxDimensions {
		if kingPosition.row+2 < maxDimensions {
			if board.fields[kingPosition.col+1][kingPosition.row+2] != nil &&
				board.fields[kingPosition.col+1][kingPosition.row+2].Piece().color != color &&
				board.fields[kingPosition.col+1][kingPosition.row+2].Type() == KnightType {
				attacked = true
			}
		}
		if kingPosition.row-2 >= 0 {
			if board.fields[kingPosition.col+1][kingPosition.row-2] != nil &&
				board.fields[kingPosition.col+1][kingPosition.row-2].Piece().color != color &&
				board.fields[kingPosition.col+1][kingPosition.row-2].Type() == KnightType {
				attacked = true
			}
		}
	}
	if kingPosition.col-1 >= 0 {
		if kingPosition.row+2 < maxDimensions {
			if board.fields[kingPosition.col-1][kingPosition.row+2] != nil &&
				board.fields[kingPosition.col-1][kingPosition.row+2].Piece().color != color &&
				board.fields[kingPosition.col-1][kingPosition.row+2].Type() == KnightType {
				attacked = true
			}
		}
		if kingPosition.row-2 >= 0 {
			if board.fields[kingPosition.col-1][kingPosition.row-2] != nil &&
				board.fields[kingPosition.col-1][kingPosition.row-2].Piece().color != color &&
				board.fields[kingPosition.col-1][kingPosition.row-2].Type() == KnightType {
				attacked = true
			}
		}
	}
	return attacked
}

func AttackedByPawn(kingPosition, freeField, blockedField Coordinates, board *Board, color Color) bool {
	attacked := false
	searchDirection := 1

	if board.direction == Reversed {
		searchDirection *= -1
	}
	if color == White {
		searchDirection *= -1
	}

	if kingPosition.row+searchDirection >= 0 && kingPosition.row+searchDirection < maxDimensions {
		if kingPosition.col-1 >= 0 {
			field := board.fields[kingPosition.col-1][kingPosition.row+searchDirection]
			if field != nil {
				if field.Type() == PawnType && field.Piece().color != color {
					attacked = true
				}
			}
		}

		if kingPosition.col+1 < maxDimensions {
			field := board.fields[kingPosition.col+1][kingPosition.row+searchDirection]
			if field != nil {
				if field.Type() == PawnType && field.Piece().color != color {
					attacked = true
				}
			}
		}
	}
	return attacked
}

func AttackedByBishopQueen(kingPosition, freeField, blockedField Coordinates, board *Board, color Color) bool {
	attacked := false
	for i := 1; kingPosition.col+i < maxDimensions && kingPosition.row+i < maxDimensions; i++ {
		if (blockedField == Coordinates{kingPosition.col + i, kingPosition.row + i}) {
			break
		}
		if (freeField == Coordinates{kingPosition.col + i, kingPosition.row + i}) {
			continue
		}
		field := board.fields[kingPosition.col+i][kingPosition.row+i]
		if field != nil {
			if field.Piece().color != color && (field.Type() == BishopType || field.Type() == QueenType) {
				attacked = true
			}
			break
		}
	}
	for i := 1; kingPosition.col-i >= 0 && kingPosition.row-i >= 0; i++ {
		if (blockedField == Coordinates{kingPosition.col - i, kingPosition.row - i}) {
			break
		}
		if (freeField == Coordinates{kingPosition.col - i, kingPosition.row - i}) {
			continue
		}
		field := board.fields[kingPosition.col-i][kingPosition.row-i]
		if field != nil {
			if field.Piece().color != color && (field.Type() == BishopType || field.Type() == QueenType) {
				attacked = true
			}
			break
		}
	}
	for i := 1; kingPosition.col+i < maxDimensions && kingPosition.row-i >= 0; i++ {
		if (blockedField == Coordinates{kingPosition.col + i, kingPosition.row - i}) {
			break
		}
		if (freeField == Coordinates{kingPosition.col + i, kingPosition.row - i}) {
			continue
		}
		field := board.fields[kingPosition.col+i][kingPosition.row-i]
		if field != nil {
			if field.Piece().color != color && (field.Type() == BishopType || field.Type() == QueenType) {
				attacked = true
			}
			break
		}
	}
	for i := 1; kingPosition.col-i >= 0 && kingPosition.row+i < maxDimensions; i++ {
		if (blockedField == Coordinates{kingPosition.col - i, kingPosition.row + i}) {
			break
		}
		if (freeField == Coordinates{kingPosition.col - i, kingPosition.row + i}) {
			continue
		}
		field := board.fields[kingPosition.col-i][kingPosition.row+i]
		if field != nil {
			if field.Piece().color != color && (field.Type() == BishopType || field.Type() == QueenType) {
				attacked = true
			}
			break
		}
	}
	return attacked
}

func AttackedByRookQueen(kingPosition, freeField, blockedField Coordinates, board *Board, color Color) bool {
	attacked := false

	for direction := 1; direction >= -1; direction -= 2 {
		for row := kingPosition.row + direction; row < maxDimensions && row >= 0; row += direction {
			if (blockedField == Coordinates{kingPosition.col, row}) {
				break
			}
			if (freeField == Coordinates{kingPosition.col, row}) {
				continue
			}
			field := board.fields[kingPosition.col][row]
			if field != nil {
				if field.Piece().color != color && (field.Type() == RookType || field.Type() == QueenType) {
					attacked = true
				}
				break
			}
		}
	}

	for direction := 1; direction >= -1; direction -= 2 {
		for col := kingPosition.col + direction; col < maxDimensions && col >= 0; col += direction {
			if (blockedField == Coordinates{col, kingPosition.row}) {
				break
			}
			if (freeField == Coordinates{col, kingPosition.row}) {
				continue
			}
			field := board.fields[col][kingPosition.row]
			if field != nil {
				if field.Piece().color != color && (field.Type() == RookType || field.Type() == QueenType) {
					attacked = true
				}
				break
			}
		}
	}

	return attacked
}

func GetMovesUp(c Coordinates, color Color, b *Board) []Coordinates {
	var fields []Coordinates
	row := c.row - 1

	for row >= 0 {
		field := b.fields[c.col][row]
		if field == nil || field.Piece().color != color {
			fields = append(fields, Coordinates{c.col, row})
		}
		if field != nil {
			break
		}
		row--
	}

	return fields
}

func GetMovesDown(c Coordinates, color Color, b *Board) []Coordinates {
	var fields []Coordinates
	row := c.row + 1

	for row < maxDimensions {
		field := b.fields[c.col][row]
		if field == nil || field.Piece().color != color {
			fields = append(fields, Coordinates{c.col, row})
		}
		if field != nil {
			break
		}
		row++
	}

	return fields
}

func GetMovesLeft(c Coordinates, color Color, b *Board) []Coordinates {
	var fields []Coordinates
	col := c.col - 1

	for col >= 0 {
		field := b.fields[col][c.row]
		if field == nil || field.Piece().color != color {
			fields = append(fields, Coordinates{col, c.row})
		}
		if field != nil {
			break
		}
		col--
	}

	return fields
}

func GetMovesRight(c Coordinates, color Color, b *Board) []Coordinates {
	var fields []Coordinates
	col := c.col + 1

	for col < maxDimensions {
		field := b.fields[col][c.row]
		if field == nil || field.Piece().color != color {
			fields = append(fields, Coordinates{col, c.row})
		}
		if field != nil {
			break
		}
		col++
	}

	return fields
}

//func validMove
