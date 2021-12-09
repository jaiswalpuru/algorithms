/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the
input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Example 4:
Input: s = "([)]"
Output: false

Example 5:
Input: s = "{[]}"
Output: true
*/

package main

import "fmt"

func is_valid(str string) bool {
	n := len(str)
	if n == 0 {
		return true
	}

	stack := make([]byte, 0)

	for i := 0; i < n; i++ {
		if str[i] == '{' || str[i] == '[' || str[i] == '(' {
			stack = append(stack, str[i])
		} else {
			fmt.Println(string(str[i]))
			m := len(stack)
			if m == 0 {
				return false
			} else {
				if str[i] == ')' {
					if stack[m-1] == '(' {
						stack = stack[:m-1]
					} else {
						return false
					}
				} else if str[i] == ']' {
					if stack[m-1] == '[' {
						stack = stack[:m-1]
					} else {
						return false
					}
				} else if str[i] == '}' {
					if stack[m-1] == '{' {
						stack = stack[:m-1]
					} else {
						return false
					}
				}
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(is_valid("()"))
	fmt.Println(is_valid("()[]{}"))
	fmt.Println(is_valid("([)]"))
}
