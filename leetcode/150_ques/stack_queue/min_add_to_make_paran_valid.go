/*
A parentheses string is valid if and only if:

It is the empty string,
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.
You are given a parentheses string s. In one move, you can insert a parenthesis at any position of the string.

For example, if s = "()))", you can insert an opening parenthesis to be "(()))" or a closing parenthesis
to be "())))".
Return the minimum number of moves required to make s valid.

Example 1:
Input: s = "())"
Output: 1

Example 2:
Input: s = "((("
Output: 3

Example 3:
Input: s = "()"
Output: 0

Example 4:
Input: s = "()))(("
Output: 4
*/

package main

import "fmt"

func min_add_to_make_valid_paran(str string) int {
	if str == "" {
		return -1
	}
	res := 0
	n := len(str)
	stack := []byte{}
	for i := 0; i < n; i++ {
		if str[i] == '(' {
			stack = append(stack, str[i])
		} else {
			m := len(stack)
			if m == 0 {
				res++
			} else {
				if stack[m-1] != '(' {
					res++
					stack = stack[:m-1]
				} else {
					stack = stack[:m-1]
				}
			}
		}
	}
	if len(stack) > 0 {
		res += len(stack)
	}
	return res
}

func main() {
	fmt.Println(min_add_to_make_valid_paran("())"))
	fmt.Println(min_add_to_make_valid_paran("((("))
	fmt.Println(min_add_to_make_valid_paran("()"))
}
