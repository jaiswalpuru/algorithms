package main

import "fmt"

func minimum_deletions_reqd_to_make_characters_unique(str string) int {
	freq := make([]int, 26)
	for i:=0; i<len(str); i++{
		freq[str[i]-'a']++
	}
	seen := make(map[int]bool)
	min_deletions := 0
	for i:=0; i<26;i++{
		for freq[i] > 0 && seen[freq[i]] {
			freq[i]--
			min_deletions++
		}
		seen[freq[i]] = true
	}
	return min_deletions
}

func main() {
	fmt.Println(minimum_deletions_reqd_to_make_characters_unique("ceabaacb"))
}
