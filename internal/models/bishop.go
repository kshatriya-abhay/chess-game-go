package models

type Bishop struct {
	position Position
	color    int
}

func MakeBishop(color, x, y int) *Bishop {
	return &Bishop{
		position: Position{Row: x, Column: y},
		color:    color,
	}
}

func (bishop *Bishop) GetPieceSymbol() string {
	if bishop.color == WHITE_COLOR {
		return "♗"
	}
	return "♝"
}

func (bishop *Bishop) GetPieceColor() int {
	return bishop.color
}

func (bishop *Bishop) GetAllowedMoves(board Board, position Position) []Position {
	return GetAllowedBishopMoves(board, position)
}

func (bishop *Bishop) IsKing(color int) bool {
	return false
}

func GetAllowedBishopMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	allowedList = append(allowedList, getMainDiagonalMoves(board, position)...)
	allowedList = append(allowedList, getAntiDiagonalMoves(board, position)...)
	return allowedList
}