package main

import (
	"fmt"
	"math"
)

func count_number_with_unique_digit(n int) int {
	visited := make([]bool, 10)
	count := 0
	backtrack(0, &count, &visited, float64(n))
	return count
}

func backtrack(curr int, count *int, visited *[]bool, n float64) {
	if curr >= int(math.Pow(10, n)) {
		fmt.Println("gere", curr)
		return
	}

	(*count)++
	fmt.Println(curr)
	for i := 0; i <= 9; i++ {
		if (i == 0 && curr == 0) || (*visited)[i] {
			continue
		}
		curr = 10*curr + i
		(*visited)[i] = true
		backtrack(curr, count, visited, n)
		curr /= 10
		(*visited)[i] = false
	}
}

func main() {
	fmt.Println(count_number_with_unique_digit(2))
}
