package main

import (
	"fmt"
	"strconv"
)

func open_the_lock(deadends []string, target string) int {
	hash_map := make(map[string]bool)
	for i := 0; i < len(deadends); i++ {
		hash_map[deadends[i]] = true
	}
	depth := 0
	q := []string{"0000"}
	visited := make(map[string]bool)
	visited["0000"] = true
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			curr := []byte(q[0])
			q = q[1:]
			if hash_map[string(curr)] {
				continue
			}
			if string(curr) == target {
				return depth
			}
			for j := 0; j < 4; j++ {
				w1 := (int(curr[j]) - int('0') - 1 + 10) % 10
				w2 := (int(curr[j]) - int('0') + 1 + 10) % 10
				str1 := string(curr[:j]) + strconv.Itoa(w1) + string(curr[j+1:])
				str2 := string(curr[:j]) + strconv.Itoa(w2) + string(curr[j+1:])
				if !visited[str1] {
					q = append(q, str1)
					visited[str1] = true
				}
				if !visited[str2] {
					q = append(q, str2)
					visited[str2] = true
				}
			}
		}
		depth++
	}
	return -1
}

func main() {
	fmt.Println(open_the_lock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
