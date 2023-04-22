package main

import (
	"fmt"
)

func break_a_palindrome(palindrome string) string {
	n := len(palindrome)
	if n == 1 {
		return ""
	}
	str := []byte(palindrome)
	for i := 0; i < n/2; i++ {
		if str[i] != 'a' {
			str[i] = 'a'
			return string(str)
		}
	}
	str[n-1] = 'b'
	return string(str)
}

func main() {
	fmt.Println(break_a_palindrome("abccba"))
}
