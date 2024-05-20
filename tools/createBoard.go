package tools

import (
	"math/rand"
)

const (
	NumSnakes  = 10
	NumLadders = 10
	BoardSize  = 100
)

func GenerateSnakesAndLadders() (map[int]int, map[int]int) {
	snakes := generateSnakes()
	ladders := generateLadders(snakes)

	return snakes, ladders
}

func generateSnakes() map[int]int {
	snakes := make(map[int]int)
	taken := make(map[int]bool)

	for len(snakes) < NumSnakes {
		start := rand.Intn(BoardSize-2) + 2
		end := rand.Intn(start-1) + 1

		if !taken[start] && !taken[end] {
			snakes[start] = end
			taken[start] = true
			taken[end] = true
		}
	}

	return snakes
}

func generateLadders(snakes map[int]int) map[int]int {
	ladders := make(map[int]int)
	taken := make(map[int]bool)

	for len(ladders) < NumLadders {
		start := rand.Intn(BoardSize-3) + 2
		end := rand.Intn(BoardSize-start-1) + start + 1

		if !taken[start] && !taken[end] && !isSnakeHead(snakes, start) && !isSnakeTail(snakes, end) {
			ladders[start] = end
			taken[start] = true
			taken[end] = true
		}
	}

	return ladders
}

func isSnakeHead(snakes map[int]int, position int) bool {
	_, ok := snakes[position]
	return ok
}

func isSnakeTail(snakes map[int]int, position int) bool {
	for _, tail := range snakes {
		if tail == position {
			return true
		}
	}
	return false
}
