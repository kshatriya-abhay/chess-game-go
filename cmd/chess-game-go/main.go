package main

import (
	"chess-game-go/pkg/chess"
	"log"
)

func main() {
	log.SetPrefix("main: ")
	log.SetFlags(0)

	log.Println("starting chess game")

	p1 := chess.CreatePlayer(1, "Abhay")
	p2 := chess.CreatePlayer(2, "Luffy")

	game := chess.CreateChessGame(p1.GetId(), p2.GetId())
	game.PrintGameBoard()

	game.MakeMoveXY(1, 1, 4, 2, 4)
	game.PrintGameBoard()

	game.MakeMoveXY(2, 6, 3, 5, 3)
	game.PrintGameBoard()

	game.MakeMoveXY(1, 2, 4, 3, 4)
	game.PrintGameBoard()
}
