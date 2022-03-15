package main

import "fmt"

func minimum_remove_to_make_valid_parantheses(str string) string {
	stack := []byte{}
	n := len(str)
	open_paran := 0
	for i := 0; i < n; i++ {
		if str[i] == ')' {
			if open_paran > 0 {
				stack = append(stack, str[i])
				open_paran--
			}
		} else {
			if str[i] == '(' {
				open_paran++
			}
			stack = append(stack, str[i])
		}
	}
	ans := ""
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == '(' && open_paran > 0 {
			open_paran--
		} else {
			ans = string(stack[i]) + ans
		}
	}
	return ans
}

func main() {
	fmt.Println(minimum_remove_to_make_valid_parantheses("())()((("))
}
