package main

import "fmt"

func alien_dictionary(words []string) string {
	graph := make(map[byte][]byte)
	cnt := make(map[byte]int)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			cnt[words[i][j]] = 0
		}
	}
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		if len(w1) > len(w2) && w1[:len(w2)] == w2 {
			return ""
		}
		for j := 0; j < min(len(w1), len(w2)); j++ {
			if w1[j] != w2[j] {
				graph[w1[j]] = append(graph[w1[j]], w2[j])
				cnt[w2[j]]++
				break
			}
		}
	}
	q := []byte{}
	for k, v := range cnt {
		if v == 0 {
			q = append(q, k)
		}
	}
	res := ""
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		res += string(curr)
		for i := 0; i < len(graph[curr]); i++ {
			cnt[graph[curr][i]]--
			if cnt[graph[curr][i]] == 0 {
				q = append(q, graph[curr][i])
			}
		}
	}
	for _, v := range cnt {
		if v != 0 {
			return ""
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(alien_dictionary([]string{"wrt", "wrf", "er", "ett", "rftt"}))
}
