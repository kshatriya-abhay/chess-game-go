package models

const WHITE_COLOR = 0
const BLACK_COLOR = 1

type Piece interface {
	GetPieceSymbol() string
	GetPieceColor() int
	GetAllowedMoves(board Board, position Position) []Position
	IsKing(color int) bool
}

func IsSameColorPiece(p1, p2 Piece) bool {
	return p1.GetPieceColor() == p2.GetPieceColor()
}

func MakePieceCopy(piece Piece, i, j int) Piece {
	switch piece.(type) {
	case *Rook:
		return MakeRook(piece.GetPieceColor(), i, j)
	case *Bishop:
		return MakeBishop(piece.GetPieceColor(), i, j)
	case *Queen:
		return MakeQueen(piece.GetPieceColor(), i, j)
	case *King:
		return MakeKing(piece.GetPieceColor(), i, j)
	case *Knight:
		return MakeKnight(piece.GetPieceColor(), i, j)
	case *Pawn:
		return MakePawn(piece.GetPieceColor(), i, j)
	}
	return nil
}
