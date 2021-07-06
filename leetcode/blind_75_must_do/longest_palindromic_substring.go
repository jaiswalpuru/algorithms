/*
Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.


Example 2:

Input: s = "cbbd"
Output: "bb"


Example 3:

Input: s = "a"
Output: "a"


Example 4:

Input: s = "ac"
Output: "a"
*/

package main

import "fmt"

func longest_palindromic_substring(str string) string {
	n := len(str)
	if n == 0 || n == 1 {
		return str
	}
	if n == 2 {
		if str[0] == str[1] {
			return str
		} else {
			return string(str[0])
		}
	}

}

func main() {
	fmt.Println(longest_palindromic_substring("babad"))
	fmt.Println(longest_palindromic_substring("cbbd"))
	fmt.Println(longest_palindromic_substring("a"))
	fmt.Println(longest_palindromic_substring("ac"))
}
