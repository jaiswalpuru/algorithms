/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

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

func valid_paran(str string) bool {
	n := len(str)
	if n == 0 || n == 1 {
		return false
	}

	i := 0
	stack := []byte{}
	for i < n {
		if str[i] == '{' || str[i] == '(' || str[i] == '[' {
			stack = append(stack, str[i])
		} else {
			if str[i] == '}' || str[i] == ')' || str[i] == ']' {

				if len(stack) == 0 {
					return false
				}
				if (str[i] == '}' && stack[len(stack)-1] != '{') || (str[i] == ')' && stack[len(stack)-1] != '(') || str[i] == ']' && stack[len(stack)-1] != '[' {
					return false
				}

				stack = stack[:len(stack)-1]
			}
		}
		i++
	}

	if i == n && len(stack) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(valid_paran("()"))
	fmt.Println(valid_paran("()[]{}"))
	fmt.Println(valid_paran("(]"))
	fmt.Println(valid_paran("([)]"))
	fmt.Println(valid_paran("{[]}"))
}
