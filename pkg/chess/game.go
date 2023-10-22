package chess

import (
	"chess-game-go/internal/models"
	"chess-game-go/internal/rules"
	"errors"
	"log"
)

type ChessGame struct {
	board         models.Board
	whitePlayer   int
	blackPlayer   int
	currentPlayer int
}

func CreateChessGame(p1, p2 int) ChessGame {
	game := ChessGame{
		board:         rules.SetupChessBoard(),
		whitePlayer:   p1,
		blackPlayer:   p2,
		currentPlayer: p1,
	}
	return game
}

func (chessGame *ChessGame) PrintGameBoard() {
	chessGame.board.PrintBoard()
}

func (chessGame *ChessGame) MakeMove(player int, startSquare, endSquare string) {
	// implement notation later
}

func (chessGame *ChessGame) MakeMoveXY(player, x1, y1, x2, y2 int) (bool, error) {
	playerColor, error := getPlayerColor(chessGame, player)
	if error != nil {
		log.Fatal(error)
	}
	if !chessGame.board.IsMoveValid(playerColor, x1, y1, x2, y2) {
		log.Fatal(errors.New("invalid move"))
	}
	result, error := chessGame.board.MovePieceOnBoard(playerColor, x1, y1, x2, y2)
	log.Println("old player", chessGame.currentPlayer)
	chessGame.updateCurrentPlayerAfterMove()
	if error != nil {
		log.Fatal(error)
	}
	log.Println("new player", chessGame.currentPlayer)
	return result, error
}

func (chessGame *ChessGame) updateCurrentPlayerAfterMove() {
	if chessGame.currentPlayer == chessGame.whitePlayer {
		chessGame.currentPlayer = chessGame.blackPlayer
	} else {
		chessGame.currentPlayer = chessGame.whitePlayer
	}
}

func getPlayerColor(chessGame *ChessGame, player int) (int, error) {
	if chessGame.currentPlayer == player && player == chessGame.whitePlayer {
		return 0, nil
	}
	if chessGame.currentPlayer == player && player == chessGame.blackPlayer {
		return 1, nil
	}
	return -1, errors.New("not this player's move")
}
