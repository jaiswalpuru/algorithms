package main

import "fmt"

//concept of connected components based on coloring the graphs

func satisfy(equation []string) bool {
	graph := make(map[int][]int)
	n := len(equation)
	for i := 0; i < n; i++ {
		if equation[i][1] == '=' {
			a, b := int(equation[i][0]-'a'), int(equation[i][3]-'a')
			graph[a] = append(graph[a], b)
			graph[b] = append(graph[b], a)
		}
	}

	color := make([]int, 26)
	c := 0

	for i := 0; i < 26; i++ {
		if color[i] == 0 {
			c++
			stack := []int{i}
			for len(stack) > 0 {
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for j := 0; j < len(graph[curr]); j++ {
					if color[graph[curr][j]] == 0 {
						color[graph[curr][j]] = c
						stack = append(stack, graph[curr][j])
					}
				}
			}
		}
	}

	for _, v := range equation {
		if v[1] == '!' {
			a, b := int(v[0]-'a'), int(v[3]-'a')
			if a == b || color[a] != 0 && color[a] == color[b] {
				return false
			}
		}
	}

	return true
}

//-------------------------------recursive----------------------------
func satisfy_recursive(equation []string) bool {
	graph := make(map[int][]int)
	n := len(equation)

	for i := 0; i < n; i++ {
		if equation[i][1] == '=' {
			a, b := int(equation[i][0]-'a'), int(equation[i][3]-'a')
			graph[a] = append(graph[a], b)
			graph[b] = append(graph[b], a)

		}
	}

	color := make([]int, 26)
	for i := 0; i < 26; i++ {
		color[i] = -1
	}
	for i := 0; i < 26; i++ {
		if color[i] == -1 {
			dfs(i, i, &color, graph)
		}
	}

	for i := 0; i < n; i++ {
		if equation[i][1] == '!' {
			a, b := int(equation[i][0]-'a'), int(equation[i][3]-'a')
			if color[a] == color[b] {
				return false
			}
		}
	}
	return true
}

func dfs(node, count int, color *[]int, graph map[int][]int) {
	if (*color)[node] == -1 {
		(*color)[node] = count
		for i := 0; i < len(graph[node]); i++ {
			dfs(graph[node][i], count, color, graph)
		}
	}
}

//-------------------------------recursive----------------------------

func main() {
	fmt.Println(satisfy_recursive([]string{"a==b", "b==a"}))
}
