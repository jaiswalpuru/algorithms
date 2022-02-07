package main

import "fmt"

func find_the_difference(s1, s2 string) byte {
	n, m := len(s1), len(s2)
	visited := make(map[byte]int)
	for i := 0; i < n; i++ {
		visited[s1[i]]++
	}

	for j := 0; j < m; j++ {
		if _, ok := visited[s2[j]]; !ok {
			return s2[j]
		} else {
			visited[s2[j]]--
			if visited[s2[j]] == 0 {
				delete(visited, s2[j])
			}
		}
	}
	return '-'
}

func main() {
	fmt.Println(find_the_difference("", ""))
}
