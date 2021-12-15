/*
Given a word w and a string s, find all the indices in s which are the starting locations of anagrams of w.
eg : w -> ab, s->abxaba return [0,3,4]
*/

package main

import "fmt"

func delete_if_zero(mp *map[rune]int, char rune) {
	if (*mp)[char] == 0 {
		delete(*mp, char)
	}
}

func anagramIndices(w, s string) []int {
	indices := []int{}
	mp := make(map[rune]int)

	for i := 0; i < len(w); i++ {
		mp[rune(w[i])]++
	}

	for i := 0; i < len(s[:len(w)]); i++ {
		mp[rune(s[i])]--
		delete_if_zero(&mp, rune(s[i]))
	}

	if len(mp) == 0 {
		indices = append(indices, 0)
	}

	for i := len(w); i < len(s); i++ {

		si, ei := s[i-len(w)], s[i] //start index , end index
		mp[rune(si)]++
		delete_if_zero(&mp, rune(si))

		mp[rune(ei)]--
		delete_if_zero(&mp, rune(ei))

		fmt.Println(string(si), string(ei), mp)
		if len(mp) == 0 {
			bi := i - len(w) + 1 //beginning index
			indices = append(indices, bi)
		}
	}

	return indices
}

func main() {
	w := "ab"
	s := "abxaba"

	fmt.Println(anagramIndices(w, s))
}
