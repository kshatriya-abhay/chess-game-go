package models

type Pawn struct {
	position Position
	color    int
}

func MakePawn(color, x, y int) *Pawn {
	return &Pawn{
		position: Position{Row: x, Column: y},
		color:    color,
	}
}

func (pawn *Pawn) GetPieceSymbol() string {
	if pawn.color == WHITE_COLOR {
		return "♙"
	}
	return "♟"
}

func (pawn *Pawn) GetPieceColor() int {
	return pawn.color
}

func (pawn *Pawn) GetAllowedMoves(board Board, position Position) []Position {
	return GetAllowedPawnMoves(board, position)
}

func (pawn *Pawn) IsKing(color int) bool {
	return false
}

func GetAllowedPawnMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	x, y := position.Row, position.Column
	dx := 1
	// dx, startx, dstartx := 1, 1, 2
	attackCells := []int {-1, 1}

	switch (board.Grid[x][y].piece.GetPieceColor()) {
	case WHITE_COLOR:
		dx = 1
		// startx = 1
		// dstartx = 2
	case BLACK_COLOR:
		dx = -1 
		// startx = BoardLength-2
		// dstartx = -2
	}
	if !board.Grid[x+dx][y].isOccupied {
		allowedList = addMoveToListIfValid(board, allowedList, x, y, dx, 0)
	}
	for i := 0; i < len(attackCells); i++ {
		cell := board.Grid[x+dx][y+attackCells[i]]
		if cell.isOccupied && cell.piece.GetPieceColor() != board.Grid[x][y].piece.GetPieceColor() {
			allowedList = addMoveToListIfValid(board, allowedList, x, y, dx, attackCells[i])
		}
	}
	// double move
	// enpassant

	return allowedList
}
