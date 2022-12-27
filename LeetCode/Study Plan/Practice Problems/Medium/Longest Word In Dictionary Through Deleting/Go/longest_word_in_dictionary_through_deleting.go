package main

import (
	"fmt"
	"sort"
)

func longest_word_in_dictionary_through_deleting(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[i]) < len(dictionary[j])
	})
	res := ""
	for i := len(dictionary) - 1; i >= 0; i-- {
		if compare_strings(s, dictionary[i]) {
			if len(res) < len(dictionary[i]) {
				res = dictionary[i]
			} else {
				if len(res) == len(dictionary[i]) {
					res = dictionary[i]
				}
			}
		}
	}
	return res
}

func compare_strings(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else {
			i++
		}
	}
	if j == len(b) {
		return true
	}
	return false
}

func main() {
	fmt.Println(longest_word_in_dictionary_through_deleting("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}
