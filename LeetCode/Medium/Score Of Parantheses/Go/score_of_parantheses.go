package main

import "fmt"

func score_of_paran_recursive(p string) int {
	return _score_of_paran_recursive(p, 0, len(p))
}

func _score_of_paran_recursive(p string, i, j int) int {
	ans, bal := 0, 0
	for k := i; k < j; k++ {
		if p[k] == '(' {
			bal += 1
		} else {
			bal -= 1
		}
		if bal == 0 {
			if k-i == 1 {
				ans += 1
			} else {
				ans += 2 * _score_of_paran_recursive(p, i+1, k)
			}
			i = k + 1
		}
	}
	return ans
}

func score_of_paran(s string) int {
	stack := []int{0}
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else {
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b+max(2*a, 1))
		}
	}
	return stack[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(score_of_paran("()((()))"))
}
