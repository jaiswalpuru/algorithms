package main

import "fmt"

func all_anagrams(s, p string) []int {
	n := len(s)
	m := len(p)

	if m > n {
		return []int{}
	}

	ans := []int{}
	mp := make(map[byte]int)
	for i := 0; i < m; i++ {
		mp[p[i]]++
	}
	visited := make(map[byte]int)

	for i := 0; i < n; i++ {
		visited[s[i]]++
		if i >= m {
			visited[s[i-m]]--
			if visited[s[i-m]] == 0 {
				delete(visited, s[i-m])
			}
		}
		if len(mp) == len(visited) {
			f := false
			for k := range visited {
				if mp[k] != visited[k] {
					f = true
				}
			}
			if !f {
				ans = append(ans, i-m+1)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(all_anagrams("cbaebabacd", "abc"))
}
