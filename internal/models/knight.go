package models

type Knight struct {
	position Position
	color    int
}

func MakeKnight(color, x, y int) *Knight {
	return &Knight{
		position: Position{Row: x, Column: y},
		color:    color,
	}
}

func (knight *Knight) GetPieceSymbol() string {
	if knight.color == WHITE_COLOR {
		return "♘"
	}
	return "♞"
}

func (knight *Knight) GetPieceColor() int {
	return knight.color
}

func (knight *Knight) GetAllowedMoves(board Board, position Position) []Position {
	return GetAllowedKnightMoves(board, position)
}

func (knight *Knight) IsKing(color int) bool {
	return false
}

func GetAllowedKnightMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	x, y := position.Row, position.Column
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 1, 2)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 2, 1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, -1, 2)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, -2, 1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, -1, -2)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, -2, -1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 1, -2)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 2, -1)
	return allowedList
}