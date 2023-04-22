/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all
non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/
package main

import (
	"fmt"
	"strings"
)

func is_valid(s string) bool {
	n := len(s)
	filter := ""
	for i := 0; i < n; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			filter += strings.ToLower(string(s[i]))
		}
		if s[i] >= '0' && s[i] <= '9' {
			filter += string(s[i])
		}
	}

	//reverse the string
	reverse_string := ""
	n = len(filter)
	for i := n - 1; i >= 0; i-- {
		reverse_string += string(filter[i])
	}
	return filter == reverse_string
}

func main() {
	fmt.Println(is_valid("A man, a plan, a canal: Panama"))
	fmt.Println(is_valid("race a car"))
	fmt.Println(is_valid("0P"))
}
