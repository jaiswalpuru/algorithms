package main

import (
	"fmt"
	"sort"
)

func longest_word_in_dictionary(dict []string) string {
	sort.Strings(dict)
	mp := make(map[string]bool)
	for i := 0; i < len(dict); i++ {
		mp[dict[i]] = true
	}
	cur_len := 0
	res := ""
	for i := len(dict) - 1; i >= 0; i-- {
		word := dict[i]
		flag := true
		for j := 1; j < len(word)+1; j++ {
			if !mp[word[:j]] {
				flag = false
				break
			}
		}
		if flag {
			if cur_len == 0 {
				cur_len = len(word)
				res = word
			} else {
				if cur_len == len(word) {
					res = word
				} else {
					if cur_len < len(word) {
						res = word
						cur_len = len(word)
					}
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(longest_word_in_dictionary([]string{"w", "wo", "wor", "worl", "world"}))
}
