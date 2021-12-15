package main

import "fmt"

type Pair struct {
	start, turn int
}

var (
	snakes = map[int]int{
		17: 13, 52: 29, 57: 40, 62: 22, 88: 18, 95: 51, 97: 79,
	}

	ladders = map[int]int{
		3: 21, 8: 30, 28: 84, 58: 77, 75: 86, 80: 100, 90: 91,
	}
)

func play() int {
	board := make(map[int]int)
	for i := 1; i <= 100; i++ {
		board[i] = i
	}

	for k, v := range snakes {
		board[k] = v
	}

	for k, v := range ladders {
		board[k] = v
	}

	start, end := 0, 100
	turns := 0

	path := []Pair{{start: start, turn: turns}}
	visited := make(map[int]bool)

	for len(path) > 0 {
		sq_turn := path[0]
		path = path[1:]

		for i := sq_turn.start + 1; i < sq_turn.start+7; i++ {
			if i >= end {
				return sq_turn.turn + 1
			}

			if _, ok := visited[i]; !ok {
				visited[i] = true
				path = append(path, Pair{start: board[i], turn: sq_turn.turn + 1})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(play())
}
