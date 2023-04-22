package main

import "math"

type WordDistance struct{ hash_map map[string][]int }

func Constructor(wordsDict []string) WordDistance {
	var d WordDistance
	hash_map := make(map[string][]int)
	n := len(wordsDict)
	for i := 0; i < n; i++ {
		hash_map[wordsDict[i]] = append(hash_map[wordsDict[i]], i)
	}
	d.hash_map = hash_map
	return d
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	l1, l2 := this.hash_map[word1], this.hash_map[word2]
	dis := math.MaxInt64
	for i := 0; i < len(l1); i++ {
		for j := 0; j < len(l2); j++ {
			dis = min(dis, abs(l1[i]-l2[j]))
		}
	}
	return dis
}

func main() {

}
