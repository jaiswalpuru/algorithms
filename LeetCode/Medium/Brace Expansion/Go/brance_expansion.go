package main

import (
	"fmt"
	"sort"
	"strings"
)

func brace_expansion(s string) []string {
	n := len(s)

	str := [][]string{}
	temp := ""
	for i := 0; i < n; i++ {
		if s[i] == '{' {
			if temp != "" {
				str = append(str, []string{temp})
				temp = ""
			}
			t := []string{}
			i++
			for i < n && s[i] != '}' {
				if s[i] != ',' {
					t = append(t, string(s[i]))
				}
				i++
			}
			str = append(str, t)
		} else {
			temp += string(s[i])
		}
	}
	if temp != "" {
		str = append(str, []string{temp})
	}

	if len(str) == 1 {
		return str[0]
	}

	res := []string{}
	_brace_expansion(str, 0, len(str), &res, []string{})
	sort.Strings(res)
	return res
}

func _brace_expansion(str [][]string, ind, n int, res *[]string, set []string) {
	if ind == n {
		*res = append(*res, strings.Join(set, ""))
		return
	}

	s := str[ind]
	for i := 0; i < len(s); i++ {
		temp := append(set, s[i])
		_brace_expansion(str, ind+1, n, res, temp)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(brace_expansion("{a,b}c{d,e}f"))
}
