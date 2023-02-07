package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {
	basket := make(map[int]int)
	size := len(fruits)

	maxFruit := 0
	k := 0

	for i := 0; i < size; i++ {
		basket[fruits[i]]++
		if len(basket) > 2 {
			for len(basket) > 2 {
				basket[fruits[k]]--
				if basket[fruits[k]] == 0 {
					delete(basket, fruits[k])
				}
				k++
			}
		}
		maxFruit = max(maxFruit, (i - k + 1))
	}

	return maxFruit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(totalFruit([]int{0, 1, 6, 6, 4, 4, 6}))
}
