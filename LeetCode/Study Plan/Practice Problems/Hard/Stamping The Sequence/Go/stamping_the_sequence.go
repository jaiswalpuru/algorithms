package main

import "fmt"

type Node struct{ made, todo map[int]int }

func stamping_the_sequence(stamp, target string) []int {
	n, m := len(stamp), len(target)
	q := []int{}
	done := make([]bool, m)
	stack := []int{}
	arr := []Node{}
	for i := 0; i <= m-n; i++ {
		made, todo := make(map[int]int), make(map[int]int)
		for j := 0; j < n; j++ {
			if target[i+j] == stamp[j] {
				made[i+j] = i + j
			} else {
				todo[i+j] = i + j
			}
		}
		arr = append(arr, Node{made, todo})
		if len(todo) == 0 {
			stack = append(stack, i)
			for j := i; j < i+n; j++ {
				if !done[j] {
					q = append(q, j)
					done[j] = true
				}
			}
		}
	}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		for j := max(0, i-n+1); j <= min(m-n, i); j++ {
			if arr[j].todo[i] != 0 {
				delete(arr[j].todo, i)
				if len(arr[j].todo) == 0 {
					stack = append(stack, j)
					for k := range arr[j].made {
						if !done[k] {
							q = append(q, k)
							done[k] = true
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(done); i++ {
		if !done[i] {
			return []int{}
		}
	}
	res := []int{}
	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stamping_the_sequence("abc", "ababc"))
}
