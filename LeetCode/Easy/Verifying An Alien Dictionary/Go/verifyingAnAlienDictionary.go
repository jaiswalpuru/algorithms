package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	k := 0
	orderNum := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderNum[order[i]] = k
		k++
	}

	size := len(words)
	for k := 1; k < size; k++ {
		w1, w2 := words[k-1], words[k]
		i, j := 0, 0
		for i < len(w1) && j < len(w2) && w1[i] == w2[j] {
			i++
			j++
		}
		if (i < len(w1) && j == len(w2)) || (i < len(w1) && j < len(w2) && orderNum[w1[i]] > orderNum[w2[j]]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}
