package main

import (
	"fmt"
	"strings"
)

func reverse_words_in_string_three(s string) string {
	str := strings.Split(s, " ")
	n := len(str)
	for i := 0; i < n; i++ {
		str[i] = reverse(str[i])
	}
	return strings.Join(str, " ")
}

func reverse(s string) string {
	sb := []byte(s)
	n := len(sb)
	for i := 0; i < n/2; i++ {
		sb[i], sb[n-1-i] = sb[n-1-i], sb[i]
	}
	return string(sb)
}

func main() {
	fmt.Println(reverse_words_in_string_three("Let's take LeetCode contest"))
}
