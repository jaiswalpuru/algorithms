package main

import (
	"fmt"
	"math"
)

func shortest_word_distance_three(words []string, word1, word2 string) int {
	res := math.MaxInt64
	ind1, ind2 := -1, -1
	for i := 0; i < len(words); i++ {
		if words[i] == word1 && words[i] == word2 {
			ind1, ind2 = i, ind1
		} else if words[i] == word1 {
			ind1 = i
		} else if words[i] == word2 {
			ind2 = i
		}
		if ind1 != -1 && ind2 != -1 {
			res = min(res, abs(ind1-ind2))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	fmt.Println(shortest_word_distance_three([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
}
