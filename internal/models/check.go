package models

import (
	"reflect"
)

func IsKingInCheck(board Board, kingColor int) bool {
	// board.PrintBoard()
	return isCheckedByRook(board, kingColor) ||
		isCheckedByBishop(board, kingColor) ||
		isCheckedByQueen(board, kingColor) ||
		isCheckedByKnight(board, kingColor) ||
		isCheckedByKing(board, kingColor)
}

func isCheckedByRook(board Board, kingColor int) bool {
	kingPosition := board.GetKingPosition(kingColor)
	kx, ky := kingPosition.Row, kingPosition.Column
	ans := checkLoop(board, kingColor, kx, ky, 1, 0, &Rook{}) || checkLoop(board, kingColor, kx, ky, 0, 1, &Rook{})
	ans = ans || checkLoop(board, kingColor, kx, ky, -1, 0, &Rook{}) || checkLoop(board, kingColor, kx, ky, 0, -1, &Rook{})
	return ans
}

func isCheckedByBishop(board Board, kingColor int) bool {
	kingPosition := board.GetKingPosition(kingColor)
	kx, ky := kingPosition.Row, kingPosition.Column
	ans := checkLoop(board, kingColor, kx, ky, 1, 1, &Bishop{}) || checkLoop(board, kingColor, kx, ky, -1, 1, &Bishop{})
	ans = ans || checkLoop(board, kingColor, kx, ky, -1, -1, &Bishop{}) || checkLoop(board, kingColor, kx, ky, 1, -1, &Bishop{})
	return ans
}

func isCheckedByQueen(board Board, kingColor int) bool {
	kingPosition := board.GetKingPosition(kingColor)
	kx, ky := kingPosition.Row, kingPosition.Column
	ans := checkLoop(board, kingColor, kx, ky, 1, 0, &Queen{}) || checkLoop(board, kingColor, kx, ky, 0, 1, &Queen{})
	ans = ans || checkLoop(board, kingColor, kx, ky, -1, 0, &Queen{}) || checkLoop(board, kingColor, kx, ky, 0, -1, &Queen{})
	ans = ans || checkLoop(board, kingColor, kx, ky, 1, 1, &Queen{}) || checkLoop(board, kingColor, kx, ky, -1, 1, &Queen{})
	ans = ans || checkLoop(board, kingColor, kx, ky, -1, -1, &Queen{}) || checkLoop(board, kingColor, kx, ky, 1, -1, &Queen{})
	return ans
}


func isCheckedByKnight(board Board, kingColor int) bool {
	kx, ky := board.GetKingPosition(kingColor).Row, board.GetKingPosition(kingColor).Column
	return isCheckedFromSquare(board, kingColor, kx, ky, 1, 2, &Knight{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, 2, 1, &Knight{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, -1, 2, &Knight{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, -2, 1, &Knight{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, -1, -2, &Knight{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, -2, -1, &Knight{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, 1, -2, &Knight{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, 2, -1, &Knight{})
}

func isCheckedByKing(board Board, kingColor int) bool {
	kx, ky := board.GetKingPosition(kingColor).Row, board.GetKingPosition(kingColor).Column
	return isCheckedFromSquare(board, kingColor, kx, ky, 0, 1, &King{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, -1, 1, &King{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, -1, 0, &King{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, -1, -1, &King{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, 0, -1, &King{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, 1, -1, &King{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, 1, 0, &King{}) ||
		isCheckedFromSquare(board, kingColor, kx, ky, 1, 1, &King{})
}

func checkLoop(board Board, kingColor, kx, ky, di, dj int, samplePiece any) bool {
	for i, j := kx, ky; i < BoardLength && i >= 0 && j < BoardLength && j >= 0; i, j = i+di, j+dj {
		if kx == i && ky == j {
			continue
		}
		cell := board.Grid[i][j]
		if !cell.isOccupied {
			continue
		}
		if cell.piece.GetPieceColor() != kingColor && reflect.TypeOf(cell.piece) == reflect.TypeOf(samplePiece) {
			return true
		} else {
			break
		}
	}
	return false
}

func isCheckedFromSquare(board Board, kingColor, kx, ky, di, dj int, samplePiece any) bool {
	i, j := kx + di, ky + dj
	return i >= 0 && i < BoardLength &&
		j >= 0 && j < BoardLength &&
		board.Grid[i][j].isOccupied &&
		board.Grid[i][j].piece.GetPieceColor() != kingColor &&
		reflect.TypeOf(board.Grid[i][j].piece) == reflect.TypeOf(samplePiece)
}