/*
Write a function that reverses a string. The input string is given as an array of characters s.

Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]


Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/

package main

import "fmt"

func reverse_string(str []string, start, end int) []string {
	n := len(str)
	if n == 0 || n == 1 { //that is length ==0
		return str
	}
	if start >= end {
		return str
	}
	str[start], str[end] = str[end], str[start]
	return reverse_string(str, start+1, end-1)
}

func main() {
	fmt.Println(reverse_string([]string{"h", "e", "l", "l", "o"}, 0, 4))
	fmt.Println(reverse_string([]string{"H", "a", "n", "n", "a", "h"}, 0, 5))
}
