package models

type Queen struct {
	position Position
	color    int
}

func MakeQueen(color, x, y int) *Queen {
	return &Queen{
		position: Position{Row: x, Column: y},
		color:    color,
	}
}

func (queen *Queen) GetPieceSymbol() string {
	if queen.color == WHITE_COLOR {
		return "♕"
	}
	return "♛"
}

func (queen *Queen) GetPieceColor() int {
	return queen.color
}

func (queen *Queen) GetAllowedMoves(board Board, position Position) []Position {
	return GetAllowedQueenMoves(board, position)
}

func (queen *Queen) IsKing(color int) bool {
	return false
}

func GetAllowedQueenMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	allowedList = append(allowedList, GetAllowedRookMoves(board, position)...)
	allowedList = append(allowedList, GetAllowedBishopMoves(board, position)...)
	return allowedList
}
