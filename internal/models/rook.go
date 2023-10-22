package models

type Rook struct {
	position Position
	color    int
}

func MakeRook(color, x, y int) *Rook {
	return &Rook{
		position: Position{Row: x, Column: y},
		color:    color,
	}
}

func (rook *Rook) GetPieceSymbol() string {
	if rook.color == WHITE_COLOR {
		return "♖"
	}
	return "♜"
}

func (rook *Rook) GetPieceColor() int {
	return rook.color
}

func (rook *Rook) GetAllowedMoves(board Board, position Position) []Position {
	return GetAllowedRookMoves(board, position)
}

func (rook *Rook) IsKing(color int) bool {
	return false
}


func GetAllowedRookMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	allowedList = append(allowedList, getRowMoves(board, position)...)
	allowedList = append(allowedList, getColumnMoves(board, position)...)
	return allowedList
}
