package main

import "fmt"

func main() {
	arr := []rune{'n', '+', '+', '-', '+'} //n represents none
	ans := []int{}
	stack := []int{}
	n := len(arr) - 1

	for i := 0; i < n; i++ {
		if arr[i+1] == '-' {
			stack = append(stack, i)
		} else {
			ans = append(ans, i)
			for len(stack) > 0 {
				ans = append(ans, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}

	stack = append(stack, n)
	for len(stack) > 0 {
		ans = append(ans, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	fmt.Println(ans)
}
