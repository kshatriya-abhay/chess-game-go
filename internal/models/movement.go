package models

import (
	"log"
)

func canPieceMoveToThisSquare(sourceCell, destinationCell Cell) (canMoveHere bool, isCellOccupied bool) {
	var canMove, isOccupied bool
	isOccupied = destinationCell.isOccupied
	canMove = !isOccupied || !IsSameColorPiece(sourceCell.piece, destinationCell.piece)
	return canMove, isOccupied
}

func getRowMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	x, y := position.Row, position.Column
	allowedList = append(allowedList, movementLoop(board, x, y, 1, 0)...)
	allowedList = append(allowedList, movementLoop(board, x, y, -1, 0)...)
	return allowedList
}

func getColumnMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	x, y := position.Row, position.Column
	allowedList = append(allowedList, movementLoop(board, x, y, 0, 1)...)
	allowedList = append(allowedList, movementLoop(board, x, y, 0, -1)...)
	return allowedList
}

func getMainDiagonalMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	x, y := position.Row, position.Column
	allowedList = append(allowedList, movementLoop(board, x, y, 1, 1)...)
	allowedList = append(allowedList, movementLoop(board, x, y, -1, -1)...)
	return allowedList
}

func getAntiDiagonalMoves(board Board, position Position) []Position {
	allowedList := make([]Position, 0)
	x, y := position.Row, position.Column
	allowedList = append(allowedList, movementLoop(board, x, y, -1, 1)...)
	allowedList = append(allowedList, movementLoop(board, x, y, 1, -1)...)
	return allowedList
}

func movementLoop(board Board, kx, ky, di, dj int) []Position {
	allowedList := make([]Position, 0)
	color := board.Grid[kx][ky].piece.GetPieceColor()
	for i, j := kx, ky; i < BoardLength && j < BoardLength && i >= 0 && j >= 0; i, j = i+di, j+dj {
		if i == kx && j == ky {
			continue
		}
		canMove, isOccupied := canPieceMoveToThisSquare(board.Grid[kx][ky], board.Grid[i][j])
		if canMove {
			tempBoard := CopyBoard(board)
			_, err := tempBoard.MovePieceOnBoard(color, kx, ky, i, j)
			if err != nil {
				log.Fatalln(err)
			}
			if !IsKingInCheck(tempBoard, color) {
				allowedList = append(allowedList, Position{i, j})
			}
		}
		if isOccupied {
			break
		}
	}
	return allowedList
}

func addMoveToListIfValid(board Board, moves []Position, kx, ky, i, j int) []Position {
	color := board.Grid[kx][ky].piece.GetPieceColor()
	if kx+i < BoardLength && kx+i >= 0 && ky+j < BoardLength && ky+j >= 0 {
		canMove, _ := canPieceMoveToThisSquare(board.Grid[kx][ky], board.Grid[kx+i][ky+j])
		if canMove {
			tempBoard := CopyBoard(board)
			can, err := tempBoard.MovePieceOnBoard(color, kx, ky, kx+i, ky+j)
			if can && err == nil && !IsKingInCheck(tempBoard, color) {
				return append(moves, Position{kx + i, ky + j})
			}
		}
	}
	return moves
}
