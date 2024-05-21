package main

import (
	"Demo/tools"
	"fmt"
	"rsc.io/quote"
	"strconv"
)

const BoardSize = 100

type Player struct {
	ID       int
	Name     string
	Position int
}

func main() {

	fmt.Println(quote.Go())

	snakes, ladders := tools.GenerateSnakesAndLadders()

	var playerCount int
	fmt.Print("Enter number of players: ")
	fmt.Scanln(&playerCount)

	players := make([]Player, playerCount)

	for i := range players {
		fmt.Printf("Enter name for Player %d: ", i+1)
		fmt.Scanln(&players[i].Name)
		players[i].ID = i + 1
		players[i].Position = 1
	}

	fmt.Println("Players:")
	for _, player := range players {
		fmt.Printf("Player %d: %s\n", player.ID, player.Name)
	}

	playGame(players, snakes, ladders)

	PrintBoard(players)
}

func playGame(players []Player, snakes, ladders map[int]int) {
	for {
		for i := range players {
			rollDice := tools.RollDice()
			fmt.Printf("%s rolled dice and got %d\n", players[i].Name, rollDice)

			players[i].Position = tools.PlayerMove(players[i].Position, rollDice, snakes, ladders, players[i].Name)

			if players[i].Position == BoardSize {
				fmt.Printf("Player %s wins!\n", players[i].Name)
				return
			}
		}
	}
}

func PrintBoard(players []Player) {
	board := [100]string{}

	for i := 0; i < 100; i++ {
		board[i] = strconv.Itoa(i + 1)
	}

	for _, player := range players {
		board[player.Position-1] = player.Name
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(board[i*10+j], "\t")
		}
		fmt.Println()
	}
}
