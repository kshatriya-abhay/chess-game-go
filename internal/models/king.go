package models

type King struct {
	position Position
	color    int
}

func MakeKing(color, x, y int) *King {
	return &King{
		position: Position{Row: x, Column: y},
		color:    color,
	}
}

func (king *King) GetPieceSymbol() string {
	if king.color == WHITE_COLOR {
		return "♔"
	}
	return "♚"
}

func (king *King) GetPieceColor() int {
	return king.color
}

func (king *King) GetAllowedMoves(board Board, position Position) []Position {
	return GetAllowedKingMoves(board, position)
}

func (king *King) IsKing(color int) bool {
	return color == king.color
}

func GetAllowedKingMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	x, y := position.Row, position.Column
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 1, 1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 0, 1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, -1, 1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, -1, 0)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, -1, -1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 0, -1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 1, -1)
	allowedList = addMoveToListIfValid(board, allowedList, x, y, 1, 0)
	return allowedList
}