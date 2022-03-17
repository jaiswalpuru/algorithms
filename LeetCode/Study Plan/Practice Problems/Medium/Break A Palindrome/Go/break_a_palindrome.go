package main

import (
	"fmt"
)

func break_a_palindrome(str string) string {
	if len(str) == 1 {
		return ""
	}

	f := false
	n := len(str)
	byte_str := []byte(str)
	if n%2 == 1 {
		for i := 0; i < n/2; i++ {
			if byte_str[i] != 'a' {
				byte_str[i] = 'a'
				f = true
				break
			}
		}
		if !f {
			for i := n/2 + 1; i < n; i++ {
				if byte_str[i] != 'a' {
					byte_str[i] = 'a'
					f = true
					break
				}
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if byte_str[i] != 'a' {
				f = true
				byte_str[i] = 'a'
				break
			}
		}
	}

	if f {
		return string(byte_str)
	}
	byte_str[len(byte_str)-1] = 'b'
	return string(byte_str)
}

func main() {
	fmt.Println(break_a_palindrome("abccba"))
}
