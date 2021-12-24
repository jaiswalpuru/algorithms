/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

package main

import "fmt"

func valid(s string) bool {

	n := len(s)
	stack := []string{}

	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, string(s[i]))
		} else {
			m := len(stack)
			if m == 0 {
				return false
			} else if s[i] == ')' && stack[m-1] == "(" {
				stack = stack[:len(stack)-1]
			} else if s[i] == '}' && stack[m-1] == "{" {
				stack = stack[:len(stack)-1]
			} else if s[i] == ']' && stack[m-1] == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(valid("()"))
	fmt.Println(valid("()[]{}"))
	fmt.Println(valid("(]"))
	fmt.Println(valid("["))
}
