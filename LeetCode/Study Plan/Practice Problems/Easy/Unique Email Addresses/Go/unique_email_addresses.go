package main

import (
	"fmt"
	"strings"
)

func unique_email_addresses(email []string) int {
	visited := make(map[string]bool)
	for i := 0; i < len(email); i++ {
		s := email[i]
		split := strings.Split(s, "@")
		split_on_plus := strings.Split(split[0], "+")
		split_on_dot := strings.Split(split_on_plus[0], ".")
		s_final := strings.Join(split_on_dot, "")
		s_final += "@" + split[1]
		visited[s_final] = true
	}
	return len(visited)
}

func main() {
	fmt.Println(unique_email_addresses([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
