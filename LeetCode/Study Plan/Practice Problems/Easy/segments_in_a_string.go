/*
Given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.
*/

package main

import (
	"fmt"
	"strings"
)

func segments(s string) int {
	if s == "" {
		return 0
	}
	l := strings.Split(s, " ")
	n := len(l)
	cnt := 0
	for i := 0; i < n; i++ {
		if strings.TrimSpace(l[i]) != "" {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(segments("Hello, my name is John"))
	fmt.Println(segments("   "))
}
