package models_test

import (
	"chess-game-go/internal/models"
	"chess-game-go/internal/util"
	"fmt"
	"testing"
)

const gotWantFailureString string = "got %d want %d"
const moveListMismatch string = "actual list %d, expected list %d"

func TestGetAllowedRookMoves(t *testing.T) {
	board := models.NewBoard()
	board.WhiteKingPosition = models.Position{0, 4}
	board.BlackKingPosition = models.Position{7, 4}
	board.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	board.Grid[3][3].OccupyCell(models.MakeRook(0, 3, 3))
	board.PrintBoard()
	moves := models.GetAllowedRookMoves(board, models.Position{3, 3}) // empty board
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 14, gotWantFailureString, len(moves), 14)
	var rookTest1 = []models.Position {
		{0, 3}, {3, 0},
		{1, 3}, {3, 1},
		{2, 3}, {3, 2},
		{4, 3}, {3, 4},
		{5, 3}, {3, 5},
		{6, 3}, {3, 6},
		{7, 3}, {3, 7},
	}
	util.AssertTrue(t, util.Compare(moves, rookTest1), moveListMismatch, moves, rookTest1)
	
	board.Grid[1][3].OccupyCell(models.MakePawn(0, 1, 3))
	board.Grid[5][3].OccupyCell(models.MakePawn(0, 5, 3))
	board.Grid[3][1].OccupyCell(models.MakePawn(0, 3, 1))
	board.Grid[3][5].OccupyCell(models.MakePawn(0, 3, 5))
	board.PrintBoard()
	moves = models.GetAllowedRookMoves(board, models.Position{3, 3}) // friendly piece in range
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 4, gotWantFailureString, len(moves), 4)
	var rookTest2 = []models.Position {
		{2, 3}, {3, 2},
		{4, 3}, {3, 4},
	}
	util.AssertTrue(t, util.Compare(moves, rookTest2), moveListMismatch, moves, rookTest2)

	board.Grid[1][3].OccupyCell(models.MakePawn(1, 1, 3))
	board.Grid[5][3].OccupyCell(models.MakePawn(1, 5, 3))
	board.Grid[3][1].OccupyCell(models.MakePawn(1, 3, 1))
	board.Grid[3][5].OccupyCell(models.MakePawn(1, 3, 5))
	board.PrintBoard()
	moves = models.GetAllowedRookMoves(board, models.Position{3, 3}) // enemy pieces in range
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 8, gotWantFailureString, len(moves), 8)
	var rookTest3 = []models.Position {
		{1, 3}, {3, 1},
		{2, 3}, {3, 2},
		{4, 3}, {3, 4},
		{5, 3}, {3, 5},
	}
	util.AssertTrue(t, util.Compare(moves, rookTest3), moveListMismatch, moves, rookTest3)

	board2 := models.NewBoard()
	board2.WhiteKingPosition = models.Position{0, 4}
	board2.BlackKingPosition = models.Position{7, 4}
	board2.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board2.Grid[1][4].OccupyCell(models.MakeRook(0, 1, 4))
	board2.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	board2.Grid[6][4].OccupyCell(models.MakeRook(1, 6, 4))
	board2.PrintBoard()
	moves = models.GetAllowedRookMoves(board2, models.Position{1, 4})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 5, gotWantFailureString, len(moves), 5)
	var rookTest4 = []models.Position {
		{2, 4},
		{3, 4},
		{4, 4},
		{5, 4},
		{6, 4},
	}
	util.AssertTrue(t, util.Compare(moves, rookTest4), moveListMismatch, moves, rookTest4)
}

func TestGetAllowedBishopMoves(t *testing.T) {
	board := models.NewBoard()
	board.WhiteKingPosition = models.Position{0, 4}
	board.BlackKingPosition = models.Position{7, 4}
	board.Grid[3][3].OccupyCell(models.MakeBishop(0, 3, 3))
	board.PrintBoard()
	moves := models.GetAllowedBishopMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 13, gotWantFailureString, len(moves), 13)
	var bishopTest1 = []models.Position {
		{0, 0}, {0, 6},
		{1, 1}, {1, 5},
		{2, 2}, {2, 4},
		{4, 4}, {4, 2},
		{5, 5}, {5, 1},
		{6, 6}, {6, 0},
		{7, 7},
	}
	util.AssertTrue(t, util.Compare(moves, bishopTest1), moveListMismatch, moves, bishopTest1)

	board.Grid[1][1].OccupyCell(models.MakePawn(0, 1, 1))
	board.Grid[5][5].OccupyCell(models.MakePawn(0, 5, 5))
	board.Grid[1][5].OccupyCell(models.MakePawn(0, 1, 5))
	board.Grid[5][1].OccupyCell(models.MakePawn(0, 5, 1))
	moves = models.GetAllowedBishopMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 4, gotWantFailureString, len(moves), 4)
	var bishopTest2 = []models.Position {
		{2, 2}, {2, 4},
		{4, 4}, {4, 2},
	}
	util.AssertTrue(t, util.Compare(moves, bishopTest2), moveListMismatch, moves, bishopTest2)

	board.Grid[1][1].OccupyCell(models.MakePawn(1, 1, 1))
	board.Grid[5][5].OccupyCell(models.MakePawn(1, 5, 5))
	board.Grid[1][5].OccupyCell(models.MakePawn(1, 1, 5))
	board.Grid[5][1].OccupyCell(models.MakePawn(1, 5, 1))
	moves = models.GetAllowedBishopMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 8, gotWantFailureString, len(moves), 8)
	var bishopTest3 = []models.Position {
		{1, 1}, {1, 5},
		{2, 2}, {2, 4},
		{4, 4}, {4, 2},
		{5, 5}, {5, 1},
	}
	util.AssertTrue(t, util.Compare(moves, bishopTest3), moveListMismatch, moves, bishopTest3)

	board2 := models.NewBoard()
	board2.WhiteKingPosition = models.Position{0, 0}
	board2.BlackKingPosition = models.Position{7, 7}
	board2.Grid[0][0].OccupyCell(models.MakeKing(0, 0, 0))
	board2.Grid[1][1].OccupyCell(models.MakeBishop(0, 1, 1))
	board2.Grid[7][7].OccupyCell(models.MakeKing(1, 7, 7))
	board2.Grid[6][6].OccupyCell(models.MakeBishop(1, 6, 6))
	board2.PrintBoard()
	moves = models.GetAllowedBishopMoves(board2, models.Position{1, 1})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 5, gotWantFailureString, len(moves), 5)
	var bishopTest4 = []models.Position {
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
	}
	util.AssertTrue(t, util.Compare(moves, bishopTest4), moveListMismatch, moves, bishopTest4)
}

func TestGetAllowedQueenMoves(t *testing.T) {
	board := models.NewBoard()
	board.WhiteKingPosition = models.Position{0, 4}
	board.BlackKingPosition = models.Position{7, 4}
	board.Grid[3][3].OccupyCell(models.MakeQueen(0, 3, 3))
	board.PrintBoard()
	moves := models.GetAllowedQueenMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 27, gotWantFailureString, len(moves), 27)
	var queenTest1 = []models.Position {
		{0, 3}, {3, 0},
		{1, 3}, {3, 1},
		{2, 3}, {3, 2},
		{4, 3}, {3, 4},
		{5, 3}, {3, 5},
		{6, 3}, {3, 6},
		{7, 3}, {3, 7},
		{0, 0}, {0, 6},
		{1, 1}, {1, 5},
		{2, 2}, {2, 4},
		{4, 4}, {4, 2},
		{5, 5}, {5, 1},
		{6, 6}, {6, 0},
		{7, 7},
	}
	util.AssertTrue(t, util.Compare(moves, queenTest1), moveListMismatch, moves, queenTest1)

	board.Grid[1][1].OccupyCell(models.MakePawn(0, 1, 1))
	board.Grid[1][3].OccupyCell(models.MakePawn(0, 1, 3))
	board.Grid[1][5].OccupyCell(models.MakePawn(0, 1, 5))
	board.Grid[3][1].OccupyCell(models.MakePawn(0, 3, 1))
	board.Grid[3][5].OccupyCell(models.MakePawn(0, 3, 5))
	board.Grid[5][1].OccupyCell(models.MakePawn(0, 5, 1))
	board.Grid[5][3].OccupyCell(models.MakePawn(0, 5, 3))
	board.Grid[5][5].OccupyCell(models.MakePawn(0, 5, 5))
	moves = models.GetAllowedQueenMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 8, gotWantFailureString, len(moves), 8)
	var queenTest2 = []models.Position {
		{2, 2}, {2, 3}, {2, 4},
		{3, 2},         {3, 4},
		{4, 2}, {4, 3}, {4, 4},
	}
	util.AssertTrue(t, util.Compare(moves, queenTest2), moveListMismatch, moves, queenTest2)

	board.Grid[1][1].OccupyCell(models.MakePawn(1, 1, 1))
	board.Grid[1][3].OccupyCell(models.MakePawn(1, 1, 3))
	board.Grid[1][5].OccupyCell(models.MakePawn(1, 1, 5))
	board.Grid[3][1].OccupyCell(models.MakePawn(1, 3, 1))
	board.Grid[3][5].OccupyCell(models.MakePawn(1, 3, 5))
	board.Grid[5][1].OccupyCell(models.MakePawn(1, 5, 1))
	board.Grid[5][3].OccupyCell(models.MakePawn(1, 5, 3))
	board.Grid[5][5].OccupyCell(models.MakePawn(1, 5, 5))
	moves = models.GetAllowedQueenMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 16, gotWantFailureString, len(moves), 16)
	var queenTest3 = []models.Position {
		{1, 1},         {1, 3},         {1, 5},
		        {2, 2}, {2, 3}, {2, 4},
		{3, 1}, {3, 2},         {3, 4}, {3, 5},
		        {4, 2}, {4, 3}, {4, 4},
		{5, 1},         {5, 3},         {5, 5},
	}
	util.AssertTrue(t, util.Compare(moves, queenTest3), moveListMismatch, moves, queenTest3)

	board2 := models.NewBoard()
	board2.WhiteKingPosition = models.Position{0, 4}
	board2.BlackKingPosition = models.Position{7, 4}
	board2.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board2.Grid[1][4].OccupyCell(models.MakeQueen(0, 1, 4))
	board2.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	board2.Grid[6][4].OccupyCell(models.MakeQueen(1, 6, 4))
	board2.PrintBoard()
	moves = models.GetAllowedQueenMoves(board2, models.Position{1, 4})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 5, gotWantFailureString, len(moves), 5)
	var queenTest4 = []models.Position {
		{2, 4},
		{3, 4},
		{4, 4},
		{5, 4},
		{6, 4},
	}
	util.AssertTrue(t, util.Compare(moves, queenTest4), moveListMismatch, moves, queenTest4)

	board3 := models.NewBoard()
	board3.WhiteKingPosition = models.Position{0, 0}
	board3.BlackKingPosition = models.Position{7, 7}
	board3.Grid[0][0].OccupyCell(models.MakeKing(0, 0, 0))
	board3.Grid[1][1].OccupyCell(models.MakeQueen(0, 1, 1))
	board3.Grid[7][7].OccupyCell(models.MakeKing(1, 7, 7))
	board3.Grid[6][6].OccupyCell(models.MakeQueen(1, 6, 6))
	board3.PrintBoard()
	moves = models.GetAllowedQueenMoves(board3, models.Position{1, 1})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 5, gotWantFailureString, len(moves), 5)
	var queenTest5 = []models.Position {
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
	}
	util.AssertTrue(t, util.Compare(moves, queenTest5), moveListMismatch, moves, queenTest5)
}

func TestGetAllowedKnightMoves(t *testing.T) {
	board := models.NewBoard()
	board.WhiteKingPosition = models.Position{0, 4}
	board.BlackKingPosition = models.Position{7, 4}
	board.Grid[3][3].OccupyCell(models.MakeKnight(0, 3, 3))
	board.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	moves := models.GetAllowedKnightMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	knightTest1 := []models.Position {
		{4, 5}, {5, 4}, {2, 5}, {1, 4}, {2, 1}, {1, 2}, {4, 1}, {5, 2},
	}
	util.AssertTrue(t, util.Compare(moves, knightTest1), moveListMismatch, moves, knightTest1)

	board.Grid[4][5].OccupyCell(models.MakePawn(0, 4, 5))
	board.Grid[5][4].OccupyCell(models.MakePawn(0, 5, 4))
	board.Grid[2][5].OccupyCell(models.MakePawn(0, 2, 5))
	board.Grid[1][4].OccupyCell(models.MakePawn(0, 1, 4))
	board.Grid[2][1].OccupyCell(models.MakePawn(0, 2, 1))
	board.Grid[1][2].OccupyCell(models.MakePawn(0, 1, 2))
	board.Grid[4][1].OccupyCell(models.MakePawn(0, 4, 1))
	board.Grid[5][2].OccupyCell(models.MakePawn(0, 5, 2))
	moves = models.GetAllowedKnightMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	knightTest2 := []models.Position {}
	util.AssertTrue(t, util.Compare(moves, knightTest2), moveListMismatch, moves, knightTest2)

	board.Grid[4][5].OccupyCell(models.MakePawn(1, 4, 5))
	board.Grid[5][4].OccupyCell(models.MakePawn(1, 5, 4))
	board.Grid[2][5].OccupyCell(models.MakePawn(1, 2, 5))
	board.Grid[1][4].OccupyCell(models.MakePawn(1, 1, 4))
	board.Grid[2][1].OccupyCell(models.MakePawn(1, 2, 1))
	board.Grid[1][2].OccupyCell(models.MakePawn(1, 1, 2))
	board.Grid[4][1].OccupyCell(models.MakePawn(1, 4, 1))
	board.Grid[5][2].OccupyCell(models.MakePawn(1, 5, 2))
	moves = models.GetAllowedKnightMoves(board, models.Position{3, 3})
	fmt.Println(moves)
	knightTest3 := []models.Position {
		{4, 5}, {5, 4}, {2, 5}, {1, 4}, {2, 1}, {1, 2}, {4, 1}, {5, 2},
	}
	util.AssertTrue(t, util.Compare(moves, knightTest3), moveListMismatch, moves, knightTest3)

	board2 := models.NewBoard()
	board2.WhiteKingPosition = models.Position{0, 4}
	board2.BlackKingPosition = models.Position{7, 4}
	board2.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board2.Grid[1][3].OccupyCell(models.MakeKnight(0, 1, 3))
	board2.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	board2.Grid[2][5].OccupyCell(models.MakeKnight(1, 2, 5))
	board2.PrintBoard()
	moves = models.GetAllowedKnightMoves(board2, models.Position{1, 3})
	fmt.Println(moves)
	util.AssertTrue(t, len(moves) == 1, gotWantFailureString, len(moves), 1)
	var knightTest4 = []models.Position {
		{2, 5},
	}
	util.AssertTrue(t, util.Compare(moves, knightTest4), moveListMismatch, moves, knightTest4)

}

func TestGetAllowedKingMoves(t *testing.T) {
	board := models.NewBoard()
	board.WhiteKingPosition = models.Position{0, 4}
	board.BlackKingPosition = models.Position{2, 4}
	board.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board.Grid[2][4].OccupyCell(models.MakeKing(1, 2, 4))
	board.Grid[2][3].OccupyCell(models.MakeRook(1, 2, 3))
	board.PrintBoard()
	moves := models.GetAllowedKingMoves(board, models.Position{0, 4})
	fmt.Println(moves)
	var kingTest1 = []models.Position {
		{0, 5},
	}
	util.AssertTrue(t, util.Compare(moves, kingTest1), moveListMismatch, moves, kingTest1)
	board.Grid[2][5].OccupyCell(models.MakeQueen(1, 2, 5))
	board.PrintBoard()
	moves = models.GetAllowedKingMoves(board, models.Position{0, 4})
	var kingTest2 = []models.Position {}
	util.AssertTrue(t, util.Compare(moves, kingTest2), moveListMismatch, moves, kingTest2)
}
