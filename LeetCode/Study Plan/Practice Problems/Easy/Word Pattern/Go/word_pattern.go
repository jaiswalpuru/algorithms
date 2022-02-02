package main

import (
	"fmt"
	"strings"
)

func word_pattern(pattern string, s string) bool {
	str := strings.Split(s, " ")
	n := len(str)
	if len(pattern) != len(str) {
		return false
	}
	mp := make(map[string]byte)
	mp_reverse := make(map[byte]string)

	for i := 0; i < n; i++ {
		v1, ok1 := mp[str[i]]
		v2, ok2 := mp_reverse[pattern[i]]

		if (ok1 && !ok2) || (!ok1 && ok2) {
			return false
		}

		if !ok1 && !ok2 {
			mp[str[i]] = pattern[i]
			mp_reverse[pattern[i]] = str[i]
		}

		if ok1 && ok2 {
			if v1 != pattern[i] || v2 != str[i] {
				return false
			}
		}

	}
	return true
}

func main() {
	fmt.Println(word_pattern("abba", "dog cat cat dog"))
}
