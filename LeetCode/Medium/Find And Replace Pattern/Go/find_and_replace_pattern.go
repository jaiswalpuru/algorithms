package main

import "fmt"

func find_and_replace_pattern(words []string, pattern string) []string {
	n := len(pattern)
	m := len(words)
	res := []string{}
	for i := 0; i < m; i++ {
		if len(words[i]) != n {
			continue
		}
		mp := make(map[byte]byte)
		bi_map := make(map[byte]byte)
		f := true
		for j := 0; j < n; j++ {
			if v, ok := mp[pattern[j]]; ok {
				if v != words[i][j] {
					f = false
					break
				}
			}
			if v, ok := bi_map[words[i][j]]; ok {
				if v != pattern[j] {
					f = false
					break
				}
			}
			mp[pattern[j]] = words[i][j]
			bi_map[words[i][j]] = pattern[j]
		}
		if f {
			res = append(res, words[i])
		}
	}
	return res
}

func main() {
	fmt.Println(find_and_replace_pattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}
