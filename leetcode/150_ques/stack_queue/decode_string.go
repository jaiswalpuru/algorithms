/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is
being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for
those repeat numbers, k. For example, there won't be input like 3a or 2[4].

Example 1:
Input: s = "3[a]2[bc]"
Output: "aaabcbc"

Example 2:
Input: s = "3[a2[c]]"
Output: "accaccacc"

Example 3:
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"

Example 4:
Input: s = "abc3[cd]xyz"
Output: "abccdcdcdxyz"
*/

package main

import (
	"fmt"
	"strconv"
)

func decode_string(str string) string {
	stack := []byte{}
	n := len(str)
	for i := 0; i < n; {
		if str[i] >= '1' && str[i] <= '9' {
			stack = append(stack, str[i])
		} else if str[i] == '[' {
			stack = append(stack, str[i])
		} else if str[i] == ']' {
			temp := ""
			j := len(stack) - 1
			for stack[j] != '[' {
				temp = string(stack[j]) + temp
				stack = stack[:j]
				j--
			}
			j--
			num := ""
			for len(stack) > 0 && stack[j] >= '0' && stack[j] <= '9' {
				num = string(stack[j]) + num
				stack = stack[:j]
				j--
			}
			val, _ := strconv.Atoi(num)
			t := ""
			for k := 1; k <= val; k++ {
				t += temp
			}
			stack = append(stack, []byte(t)...)
		} else {
			stack = append(stack, str[i])
		}
		i++
	}
	return string(stack)
}

func main() {
	// fmt.Println(decode_string("3[a]2[bc]"))
	// fmt.Println(decode_string("abc3[cd]xyz"))
	// fmt.Println(decode_string("3[a2[c]]"))
	fmt.Println(decode_string("100[leetcode]"))
}
