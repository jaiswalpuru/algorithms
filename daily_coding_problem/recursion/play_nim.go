/*
The game of Nim is played as follows. Starting with three heaps, each containing a variable number of items,
two players take turns removing one or more items from a single pile. The player who eventually is forced to take
the last stone loses. For a single pile. The player who eventually is forced to take the last stone loses.

Given a list of non zero starting values[a,b,c] and assuming optimal play, determine whether the first player has
a forced win.
*/

package main

import "fmt"

func update(heaps []int, pile, items int) []int {
	temp := make([]int, len(heaps))
	for i := 0; i < len(heaps); i++ {
		temp[i] = heaps[i]
	}
	temp[pile] -= items
	return temp
}

func get_moves(heaps []int) [][]int {
	moves := [][]int{}
	for piles, count := range heaps {
		for i := 1; i < count+1; i++ {
			moves = append(moves, update(heaps, piles, i))
		}
	}
	// moves_set := [][]int{}
	// for i := 0; i < len(moves); i++ {
	// 	for j := 0; j < len(moves) && j != i; j++ {
	// 		if moves[i][0] == moves[j][0] && moves[i][1] == moves[j][1] && moves[i][2] == moves[j][2] {
	// 			moves[j][0] = -9
	// 			moves[j][1] = -9
	// 			moves[j][2] = -9
	// 		}
	// 	}
	// }

	// for i := 0; i < len(moves); i++ {
	// 	if moves[i][0] != -9 && moves[i][1] != -9 && moves[i][2] != -9 {
	// 		moves_set = append(moves_set, moves[i])
	// 	}
	// }

	return moves
}

func nim(heaps []int) bool {
	if heaps[0] == 0 && heaps[1] == 0 && heaps[2] == 0 {
		return true
	}

	moves := get_moves(heaps)
	res := []bool{}
	for i := range moves {
		if nim(moves[i]) != true {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	//fmt.Println(res)
	for i := 0; i < len(res); i++ {
		if res[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(nim([]int{3, 4, 5}))
	fmt.Println(nim([]int{1, 3, 0}))
}
