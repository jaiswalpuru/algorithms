/*
Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
*/

package main

import (
	"fmt"
	"strings"
)

func is_valid(str string) bool {
	pure_str := ""
	n := len(str)
	for i := 0; i < n; i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			pure_str += strings.ToLower(string(str[i]))
		} else if str[i] >= '0' && str[i] <= '9' {
			pure_str += string(str[i])
		}
	}
	return is_palindrome(pure_str)
}

func is_palindrome(str string) bool {
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(is_valid("A man, a plan, a canal: Panama"))
	fmt.Println(is_valid("race a car"))
	fmt.Println(is_valid("0P"))
}
