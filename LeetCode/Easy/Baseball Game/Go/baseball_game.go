package main

import (
	"fmt"
	"strconv"
)

func baseball_game(op []string) int {
	stack := []int{}
	n := len(op)
	for i := 0; i < n; i++ {
		if op[i] == "C" {
			stack = stack[:len(stack)-1]
		}else if op[i] == "D" {
			v := stack[len(stack)-1]
			stack = append(stack, 2*v)
		}else if op[i] == "+" {
			v1, v2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = append(stack, v1+v2)
		}else {
			v, _ := strconv.Atoi(op[i])
			stack = append(stack, v)
		}
	}

	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}

	return sum
}

func main() {
	fmt.Println(baseball_game([]string{"5", "2", "C", "D", "+"}))
}
