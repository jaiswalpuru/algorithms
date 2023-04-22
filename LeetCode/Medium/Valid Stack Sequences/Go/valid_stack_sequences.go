package main

import "fmt"

func valid_stack_sequence(pushed []int, popped []int) bool {
	stack := []int{}
	i, j, n := 0, 0, len(pushed)
	for i < n && j < n {
		if pushed[i] != popped[j] {
			f := true
			if len(stack) > 0 {
				if stack[len(stack)-1] == popped[j] {
					stack = stack[:len(stack)-1]
					i--
					j++
					f = false
				}
			}
			if f {
				stack = append(stack, pushed[i])
			}
			i++
		} else {
			i++
			j++
		}
	}

	for len(stack) > 0 && j < n {
		if stack[len(stack)-1] != popped[j] {
			return false
		}
		stack = stack[:len(stack)-1]
		j++
	}

	return true
}

func main() {
	fmt.Println(valid_stack_sequence([]int{3, 1, 0, 2}, []int{2, 0, 3, 1}))
}
