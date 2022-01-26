package main

import (
	"fmt"
	"math"
	"sort"
)

func get_val(n int) int {
	moves := 0
	for n > 0 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		moves++
		if n == 1 {
			break
		}
	}
	return moves
}

//get power value using dp
func get_power_value(n int, dp *[]int) int {
	if (*dp)[n] != -1 {
		return (*dp)[n]
	}

	if n%2 == 0 {
		(*dp)[n] = 1 + get_power_value(n/2, dp)
	} else {
		(*dp)[n] = 1 + get_power_value(n*3+1, dp)
	}

	return (*dp)[n]
}

type Pair struct {
	n,
	moves int
}

func sort_integers(lo, hi, k int) int {
	moves := []Pair{}
	dp := make([]int, math.MaxInt32)
	dp[1] = 0
	for i := lo; i <= hi; i++ {
		moves = append(moves, Pair{n: i, moves: get_power_value(i, &dp)})
	}

	sort.Slice(moves, func(i, j int) bool {
		if moves[i].moves == moves[j].moves {
			return moves[i].n < moves[j].n
		}
		return moves[i].moves < moves[j].moves
	})
	return moves[k-1].n
}

func main() {
	fmt.Println(sort_integers(1, 1000, 777))
}
