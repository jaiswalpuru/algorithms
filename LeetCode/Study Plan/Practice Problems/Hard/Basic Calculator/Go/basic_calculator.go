package main

import (
	"fmt"
	"math"
	"strconv"
)

func basic_calculator(s string) int {
	stack := []string{}
	op := 0
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			op = int(math.Pow(10, float64(n)))*int(ch-'0') + op
			n++
		} else if ch != ' ' {
			if n != 0 {
				stack = append(stack, strconv.Itoa(op))
				op = 0
				n = 0
			}
			if ch == '(' {
				res := eval(&stack)
				stack = stack[:len(stack)-1]
				stack = append(stack, strconv.Itoa(res))
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}
	if n != 0 {
		stack = append(stack, strconv.Itoa(op))
	}
	return eval(&stack)
}

func eval(stack *[]string) int {
	n := len(*stack)
	f := false
	_, err := strconv.Atoi((*stack)[n-1])
	if err != nil {
		f = true
	}
	if n == 0 || f {
		*stack = append(*stack, "0")
	}
	res, _ := strconv.Atoi((*stack)[len(*stack)-1])
	*stack = (*stack)[:len(*stack)-1]
	for len(*stack) > 0 && (*stack)[len(*stack)-1] != ")" {
		sign := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		if sign == "+" {
			v, _ := strconv.Atoi((*stack)[len(*stack)-1])
			res += v
		} else {
			v, _ := strconv.Atoi((*stack)[len(*stack)-1])
			res -= v
		}
		*stack = (*stack)[:len(*stack)-1]
	}
	return res
}

func main() {
	fmt.Println(basic_calculator("1 + 1"))
}
