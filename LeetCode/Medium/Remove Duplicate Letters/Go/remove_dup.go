package main

import (
	"fmt"
)

func remove_duplicates(str string) string {
	//mp will contain the last occurrence
	mp := make(map[byte]int)
	n := len(str)
	for i := 0; i < n; i++ {
		mp[str[i]] = i
	}

	visited := make(map[byte]struct{})
	stack := []byte{}
	for i := 0; i < n; i++ {
		curr := str[i]
		if _, ok := visited[curr]; !ok {
			for len(stack) > 0 && stack[len(stack)-1] > curr && mp[stack[len(stack)-1]] > i {
				pop := stack[len(stack)-1]
				delete(visited, pop)
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, curr)
			visited[curr] = struct{}{}
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(remove_duplicates("cdadabcc"))
}
