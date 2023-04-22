package main

import (
	"fmt"
	"reflect"
)

func count_pairs_of_similar_string(words []string) int {
	mp := make(map[string]map[byte]bool)
	for i := 0; i < len(words); i++ {
		mp[words[i]] = make(map[byte]bool)
		for j := 0; j < len(words[i]); j++ {
			mp[words[i]][words[i][j]] = true
		}
	}
	cnt := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			v1, v2 := mp[words[i]], mp[words[j]]
			if reflect.DeepEqual(v1, v2) {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(count_pairs_of_similar_string([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
}
