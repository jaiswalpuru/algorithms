package main

import (
	"fmt"
	"strconv"
)

func evaluate_reverse_polish_notation(s []string) int {
	st := []int{}
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == "+" || s[i] == "*" || s[i] == "/" || s[i] == "-" {
			v1, v2 := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, evaluate(s[i], v1, v2))
		} else {
			v, _ := strconv.Atoi(s[i])
			st = append(st, v)
		}
	}
	return st[0]
}

func evaluate(char string, v1, v2 int) int {
	if char == "+" {
		return v2 + v1
	}
	if char == "-" {
		return v2 - v1
	}
	if char == "/" {
		return v2 / v1
	}
	if char == "*" {
		return v2 * v1
	}
	return 1
}

func main() {
	fmt.Println(evaluate_reverse_polish_notation([]string{"2", "1", "+", "3", "*"}))
}
