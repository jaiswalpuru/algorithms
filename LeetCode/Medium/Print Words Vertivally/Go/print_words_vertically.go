package main

import (
	"fmt"
	"strings"
)

func print_words_vertically(s string) []string {
	res := []string{}
	str := strings.Split(s, " ")
	n := len(str)
	m := len(str[0])
	for i := 0; i < n; i++ {
		m = max(m, len(str[i]))
	}
	for j := 0; j < m; j++ {
		s := ""
		for i := 0; i < n; i++ {
			if len(str[i]) <= j {
				s += " "
				continue
			}
			s += string(str[i][j])
		}
		s = strings.TrimRight(s, " ")
		res = append(res, s)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(print_words_vertically("HOW ARE YOU"))
}
