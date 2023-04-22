package main

import "fmt"

var ascii = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
	'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func word_ladder(start, end string, list []string) int {
	mp := make(map[string]bool)
	n := len(list)
	for i := 0; i < n; i++ {
		mp[list[i]] = true
	}
	if !mp[end] {
		return 0
	}
	q := []string{start}
	l := 1
	visited := make(map[string]bool)
	visited[start] = true
	for len(q) > 0 {
		q_len := len(q)
		for m := 0; m < q_len; m++ {
			curr := q[0]
			q = q[1:]
			for i := 0; i < len(curr); i++ {
				for k := 0; k < 26; k++ {
					next_word := curr[:i] + string(ascii[k]) + curr[i+1:]
					if !mp[next_word] || curr == next_word {
						continue
					}
					if next_word == end {
						return l + 1
					}
					if !visited[next_word] {
						q = append(q, next_word)
						visited[next_word] = true
					}
				}
			}
		}
		l++
	}
	return 0
}

func main() {
	fmt.Println(word_ladder("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
