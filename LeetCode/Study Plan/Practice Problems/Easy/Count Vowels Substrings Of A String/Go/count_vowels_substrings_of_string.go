package main

import "fmt"

func count_vowels_substrings_of_string(s string) int {
	res := 0
	n := len(s)
	for i := 0; i < n; i++ {
		mp := make(map[byte]bool)
		for j := i; j < n; j++ {
			if is_vowel(s[j]) {
				mp[s[j]] = true
				if len(mp) == 5 {
					res++
				}
			} else {
				break
			}
		}
	}
	return res
}

func is_vowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func main() {
	fmt.Println(count_vowels_substrings_of_string("aeiouu"))
}
