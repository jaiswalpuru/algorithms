package main

import (
	"fmt"
)

func solve(str string) int {
	if str == "" {
		return 0
	}

	stack := []int{}
	curr_number := 0
	n := len(str)
	op := byte('+')
	for i := 0; i < n; i++ {
		c := str[i]
		if c >= '0' && c <= '9' {
			curr_number = (curr_number * 10) + int(c-'0')
		}

		if !(c >= '0' && c <= '9') && c != ' ' || i == n-1 {
			if op == '+' {
				stack = append(stack, curr_number)
			} else if op == '-' {
				stack = append(stack, -curr_number)
			} else if op == '*' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res := v * curr_number
				stack = append(stack, res)
			} else if op == '/' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res := v / curr_number
				stack = append(stack, res)
			}
			op = c
			curr_number = 0
		}
	}
	res := 0
	for len(stack) > 0 {
		res += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

func main() {
	fmt.Println(solve("1"))
}
