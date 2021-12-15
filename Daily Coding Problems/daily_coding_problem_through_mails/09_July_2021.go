/*
Given a string of parentheses, write a function to compute the minimum number of parentheses to be r
emoved to make the string valid (i.e. each open parenthesis is eventually closed).

For example, given the string "()())()", you should return 1. Given the string ")(", you should return 2,
since we must remove all of them.
*/

package main

import "fmt"

func remove(str string) int {
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
	fmt.Println(remove("()())()"))
	fmt.Println(remove(")("))
}
