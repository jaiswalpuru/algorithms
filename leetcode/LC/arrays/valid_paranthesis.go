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


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package main

import "fmt"

func is_valid(str string) bool {
	rune_arr := []rune(str)
	n := len(rune_arr)
	stack := []rune{rune_arr[0]}
	for i := 1; i < n; i++ {
		if rune_arr[i] == '}' || rune_arr[i] == ')' || rune_arr[i] == ']' {
			if len(stack) <= 0 {
				return false
			} else {
				if stack[len(stack)-1] != rune_arr[i]-2 {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		} else {
			stack = append(stack, rune_arr[i])
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(is_valid("[]"))
	fmt.Println(is_valid("(]"))
	fmt.Println(is_valid("([)]"))
	fmt.Println(is_valid("{[]}"))
}
