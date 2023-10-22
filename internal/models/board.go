package models

import (
	"chess-game-go/internal/util"
	"errors"
	"fmt"
	"log"
	"strings"
)

const BoardLength = 8

type Cell struct {
	piece      Piece
	isOccupied bool
}

func (cell *Cell) OccupyCell(piece Piece) {
	cell.piece = piece
	cell.isOccupied = true
}

func (cell *Cell) RemovePiece() {
	cell.piece = nil
	cell.isOccupied = false
}

type Board struct {
	Grid [][]Cell
	WhiteKingPosition Position
	BlackKingPosition Position
}

func NewBoard() Board {
	board := Board{
		Grid: make([][]Cell, BoardLength),
	}
	for i := range board.Grid {
		board.Grid[i] = make([]Cell, BoardLength)
	}
	return board
}

func CopyBoard(board Board) Board {
	copyBoard := NewBoard()
	copyBoard.WhiteKingPosition = Position{Row: board.WhiteKingPosition.Row, Column: board.WhiteKingPosition.Column}
	copyBoard.BlackKingPosition = Position{Row: board.BlackKingPosition.Row, Column: board.BlackKingPosition.Column}
	
	for i := 0; i < BoardLength; i++ {
		for j := 0; j < BoardLength; j++ {
			copyBoard.Grid[i][j] = Cell{
				piece: MakePieceCopy(board.Grid[i][j].piece, i, j),
				isOccupied: board.Grid[i][j].isOccupied,
			}
		}
	}
	return copyBoard
}

func (board Board) PrintBoard() {
	log.Println("Board:")
	for i := len(board.Grid) - 1; i >= 0; i-- {
		builder := strings.Builder{}
		for _, cell := range board.Grid[i] {
			if !cell.isOccupied {
				builder.WriteString(". ")
			} else {
				builder.WriteString(cell.piece.GetPieceSymbol())
				builder.WriteString(" ")
			}
		}
		log.Println(builder.String())
	}
	log.Println("----------------")
}

func (board *Board) GetKingPosition(color int) Position {
	if color == WHITE_COLOR {
		return board.WhiteKingPosition
	}
	if color == BLACK_COLOR {
		return board.BlackKingPosition
	}
	return Position{}
}

func (board *Board) UpdateKingPosition(color int, x, y int) {
	if color == WHITE_COLOR {
		board.WhiteKingPosition.Row = x
		board.WhiteKingPosition.Column = y
	}
	if color == BLACK_COLOR {
		board.BlackKingPosition.Row = x
		board.BlackKingPosition.Column = y
	}
}

func (board *Board) MovePieceOnBoard(playerColor, x1, y1, x2, y2 int) (bool, error) {
	if x1 < 0 || x1 >= BoardLength || y1 < 0 || y1 >= BoardLength {
		return false, errors.New("invalid start location")
	}
	if x2 < 0 || x2 >= BoardLength || y2 < 0 || y2 >= BoardLength {
		return false, errors.New("invalid end location")
	}
	if board.Grid[x1][y1].piece == nil {
		return false, fmt.Errorf("no piece on start square %d, %d", x1, y1)
	}
	if board.Grid[x1][y1].piece.GetPieceColor() != playerColor {
		return false, fmt.Errorf("not this player's piece on start square %d, %d", x1, y1)
	}
	if board.Grid[x2][y2].isOccupied && board.Grid[x2][y2].piece.GetPieceColor() == playerColor {
		return false, errors.New("Cell is already occupied by this player's color")
	}
	board.Grid[x2][y2].piece = board.Grid[x1][y1].piece
	board.Grid[x2][y2].isOccupied = true
	board.Grid[x1][y1].piece = nil
	board.Grid[x1][y1].isOccupied = false
	if board.Grid[x2][y2].piece.IsKing(playerColor) {
		board.UpdateKingPosition(playerColor, x2, y2)
	}
	return true, nil
}

func (board *Board) IsMoveValid(playerColor, x1, y1, x2, y2 int) bool {
	startPosition := Position{x1, y1}
	endPosition := Position{x2, y2}
	moves := board.Grid[x1][y1].piece.GetAllowedMoves(*board, startPosition)
	return util.SliceContains(moves, endPosition)
}