package main

import "fmt"

func remove(str string, k int) string {
	n := len(str)
	stack := []byte{}
	for i := 0; i < n; i++ {
		curr := str[i]
		for len(stack) > 0 && stack[len(stack)-1] > curr && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, str[i])
	}

	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}

	for k > 0 {
		stack = stack[:len(stack)-1]
		if string(stack) == "" {
			return "0"
		}
		k--
	}

	return string(stack)
}

func main() {
	fmt.Println(remove("112", 1))
}
