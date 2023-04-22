package main

import "fmt"

func valid_paran(str string) bool {
	stack := []byte{}
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] == '(' || str[i] == '{' || str[i] == '[' {
			stack = append(stack, str[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if (str[i] == ')' && stack[len(stack)-1] != '(') || (str[i] == '}' && stack[len(stack)-1] != '{') || (str[i] == ']' && stack[len(stack)-1] != '[') {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
func main() {
	fmt.Println(valid_paran("()"))
}
