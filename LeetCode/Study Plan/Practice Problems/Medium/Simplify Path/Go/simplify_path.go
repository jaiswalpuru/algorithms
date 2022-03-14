package main

import (
	"fmt"
	"strings"
)

func simplify_path(path string) string {
	s := strings.Split(path, "/")
	n := len(s)
	s_r := []string{}
	for i := 0; i < n; i++ {
		if s[i] != "" && s[i] != "." {
			s_r = append(s_r, s[i])
		}
	}
	st := []string{}
	n = len(s_r)
	for i := 0; i < n; i++ {
		if s_r[i] == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, s_r[i])
		}
	}
	return "/" + strings.Join(st, "/")
}

func main() {
	fmt.Println(simplify_path("/a/./b/../../c/"))
}
