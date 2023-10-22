package rules

import "chess-game-go/internal/models"

func SetupChessBoard() models.Board {
	board := models.NewBoard()
	board.Grid[0][0].OccupyCell(models.MakeRook(0, 0, 0))
	board.Grid[0][1].OccupyCell(models.MakeKnight(0, 0, 1))
	board.Grid[0][2].OccupyCell(models.MakeBishop(0, 0, 2))
	board.Grid[0][3].OccupyCell(models.MakeQueen(0, 0, 3))
	board.Grid[0][4].OccupyCell(models.MakeKing(0, 0, 4))
	board.Grid[0][5].OccupyCell(models.MakeBishop(0, 0, 5))
	board.Grid[0][6].OccupyCell(models.MakeKnight(0, 0, 6))
	board.Grid[0][7].OccupyCell(models.MakeRook(0, 0, 7))
	board.Grid[1][0].OccupyCell(models.MakePawn(0, 1, 0))
	board.Grid[1][1].OccupyCell(models.MakePawn(0, 1, 1))
	board.Grid[1][2].OccupyCell(models.MakePawn(0, 1, 2))
	board.Grid[1][3].OccupyCell(models.MakePawn(0, 1, 3))
	board.Grid[1][4].OccupyCell(models.MakePawn(0, 1, 4))
	board.Grid[1][5].OccupyCell(models.MakePawn(0, 1, 5))
	board.Grid[1][6].OccupyCell(models.MakePawn(0, 1, 6))
	board.Grid[1][7].OccupyCell(models.MakePawn(0, 1, 7))

	board.Grid[7][0].OccupyCell(models.MakeRook(1, 7, 0))
	board.Grid[7][1].OccupyCell(models.MakeKnight(1, 7, 1))
	board.Grid[7][2].OccupyCell(models.MakeBishop(1, 7, 2))
	board.Grid[7][3].OccupyCell(models.MakeQueen(1, 7, 3))
	board.Grid[7][4].OccupyCell(models.MakeKing(1, 7, 4))
	board.Grid[7][5].OccupyCell(models.MakeBishop(1, 7, 5))
	board.Grid[7][6].OccupyCell(models.MakeKnight(1, 7, 6))
	board.Grid[7][7].OccupyCell(models.MakeRook(1, 7, 7))
	board.Grid[6][0].OccupyCell(models.MakePawn(1, 6, 0))
	board.Grid[6][1].OccupyCell(models.MakePawn(1, 6, 1))
	board.Grid[6][2].OccupyCell(models.MakePawn(1, 6, 2))
	board.Grid[6][3].OccupyCell(models.MakePawn(1, 6, 3))
	board.Grid[6][4].OccupyCell(models.MakePawn(1, 6, 4))
	board.Grid[6][5].OccupyCell(models.MakePawn(1, 6, 5))
	board.Grid[6][6].OccupyCell(models.MakePawn(1, 6, 6))
	board.Grid[6][7].OccupyCell(models.MakePawn(1, 6, 7))

	board.WhiteKingPosition = models.Position{Row: 0, Column: 4}
	board.BlackKingPosition = models.Position{Row: 7, Column: 4}
	
	return board
}