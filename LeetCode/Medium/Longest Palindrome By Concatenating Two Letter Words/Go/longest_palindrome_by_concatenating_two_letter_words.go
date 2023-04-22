package main

import "fmt"

func longest_palindrome_by_concatenating_two_letter_words(words []string) int {
	mp := make(map[string]int)
	for i := 0; i < len(words); i++ {
		mp[words[i]]++
	}
	odd := 0
	l := 0
	val := ""
	for k, v := range mp {
		if k[0] == k[1] && v%2 == 1 {
			if v > odd {
				val = k
				odd = v
			}
		}
	}
	l += odd * 2
	mp[val] = 0
	for i := 0; i < len(words); i++ {
		k := words[i]
		if mp[k] > 0 {
			if k[0] == k[1] {
				if mp[k]%2 == 0 {
					l += mp[k] * 2
				} else {
					l += (mp[k] - 1) * 2
				}
				mp[k] = 0
			} else {
				l += (min(mp[k], mp[string(k[1])+string(k[0])]) * 4)
				mp[k] = 0
				mp[string(k[1])+string(k[0])] = 0
			}
		}
	}
	return l
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(longest_palindrome_by_concatenating_two_letter_words([]string{"lc", "cl", "gg"}))
}
