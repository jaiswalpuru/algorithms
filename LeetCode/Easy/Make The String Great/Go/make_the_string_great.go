package main

import (
	"fmt"
	"strings"
)

func make_the_string_great(s string) string {
	stack := []byte{}
	n := len(s)
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			top := stack[len(stack)-1]
			curr := s[i]
			if strings.ToLower(string(top)) == strings.ToLower(string(curr)) {
				if int(curr)-int(top) != 0 {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, s[i])
				}
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(make_the_string_great("leEeetcode"))
}
