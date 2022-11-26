package main

import (
	"fmt"
	"math"
)

func shortest_word_distance(words_dict []string, word1, word2 string) int {
	ind1, ind2 := -1, -1
	dis := math.MaxInt64
	n := len(words_dict)
	for i := 0; i < n; i++ {
		if words_dict[i] == word1 {
			ind1 = i
		}
		if words_dict[i] == word2 {
			ind2 = i
		}
		if ind1 != -1 && ind2 != -1 {
			dis = min(dis, abs(ind1-ind2))
		}
	}
	return dis
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(shortest_word_distance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
}
