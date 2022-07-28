package main

import "fmt"

func valid_anagram(s, t string) bool {
	mp_s := make(map[byte]int)
	mp_t := make(map[byte]int)
	n := len(s)
	m := len(t)
	for i := 0; i < n; i++ {
		mp_s[s[i]]++
	}
	for j := 0; j < m; j++ {
		mp_t[t[j]]++
	}
	for k, v := range mp_s {
		if mp_t[k] != v {
			return false
		}
	}
	for k, v := range mp_t {
		if mp_s[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(valid_anagram("anagram", "nagaram"))
}
