package models_test

import (
	"chess-game-go/internal/models"
	"fmt"
	"testing"
)

func TestIsKingInCheck(t *testing.T) {
	board := models.NewBoard()
	board.Grid[3][3].OccupyCell(models.MakeKing(0, 3, 3))
	board.Grid[7][7].OccupyCell(models.MakeKing(1, 7, 7))
	board.Grid[3][4].OccupyCell(models.MakeRook(1, 3, 4))
	board.Grid[7][6].OccupyCell(models.MakeRook(0, 7, 6))
	board.WhiteKingPosition = models.Position{3, 3}
	board.BlackKingPosition = models.Position{7, 7}
	fmt.Println(models.IsKingInCheck(board, models.WHITE_COLOR))
	fmt.Println(models.IsKingInCheck(board, models.BLACK_COLOR))
	
	board.Grid[3][4].RemovePiece()
	board.Grid[7][6].RemovePiece()
	board.Grid[4][2].OccupyCell(models.MakeBishop(1, 4, 2))
	board.Grid[6][6].OccupyCell(models.MakeBishop(0, 6, 6))
	fmt.Println(models.IsKingInCheck(board, models.WHITE_COLOR))
	fmt.Println(models.IsKingInCheck(board, models.BLACK_COLOR))

	board2 := models.NewBoard()
	board2.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board2.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	board2.Grid[1][5].OccupyCell(models.MakeRook(0, 1, 5))
	board2.Grid[6][4].OccupyCell(models.MakeRook(1, 6, 4))
	board2.WhiteKingPosition = models.Position{0, 4}
	board2.BlackKingPosition = models.Position{7, 4}
	fmt.Println(models.IsKingInCheck(board2, models.WHITE_COLOR))
	fmt.Println(models.IsKingInCheck(board2, models.BLACK_COLOR))

	board3 := models.NewBoard()
	board3.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board3.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	board3.Grid[1][4].OccupyCell(models.MakeRook(0, 1, 4))
	board3.Grid[6][4].OccupyCell(models.MakeRook(1, 6, 4))
	board3.WhiteKingPosition = models.Position{0, 4}
	board3.BlackKingPosition = models.Position{7, 4}
	fmt.Println(models.IsKingInCheck(board3, models.WHITE_COLOR))
	fmt.Println(models.IsKingInCheck(board3, models.BLACK_COLOR))

	board4 := models.NewBoard()
	board4.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board4.Grid[4][4].OccupyCell(models.MakeKing(1, 4, 4))
	board4.Grid[2][5].OccupyCell(models.MakeKnight(0, 2, 5))
	board4.Grid[2][3].OccupyCell(models.MakeKnight(1, 2, 3))
	board4.WhiteKingPosition = models.Position{0, 4}
	board4.BlackKingPosition = models.Position{4, 4}
	fmt.Println(models.IsKingInCheck(board4, models.WHITE_COLOR))
	fmt.Println(models.IsKingInCheck(board4, models.BLACK_COLOR))
}