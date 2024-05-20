package tools

import (
	"fmt"
	"math/rand"
)

func RollDice() int {
	return rand.Intn(6) + 1
}

func PlayerMove(position, rollDice int, snakes, ladders map[int]int, name string) int {
	newPosition := position + rollDice

	if newPosition > BoardSize {
		return position
	}

	if tail, isSnake := snakes[newPosition]; isSnake {
		fmt.Printf("%s encountered a snake! Sliding down from %d to %d\n", name, newPosition, tail)
		return tail
	}

	if end, isLadder := ladders[newPosition]; isLadder {
		fmt.Printf("%s encountered a ladder! Climbing up from %d to %d\n", name, newPosition, end)
		return end
	}

	fmt.Printf("%s Climbing up from %d to %d\n", name, newPosition, position)
	return newPosition
}
