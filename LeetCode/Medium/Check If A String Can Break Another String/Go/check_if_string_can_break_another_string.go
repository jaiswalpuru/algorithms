package main

import (
	"fmt"
	"sort"
)

func check_if_string_can_break_another_string(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1_b, s2_b := []byte(s1), []byte(s2)
	sort.Slice(s1_b, func(i, j int) bool {
		return s1_b[i] < s1_b[j]
	})
	sort.Slice(s2_b, func(i, j int) bool {
		return s2_b[i] < s2_b[j]
	})
	f, t := true, true
	for i := 0; i < len(s1_b); i++ {
		if !(s2_b[i] >= s1_b[i]) {
			f = false
			break
		}
	}
	for i := 0; i < len(s1_b); i++ {
		if !(s1_b[i] >= s2_b[i]) {
			t = false
			break
		}
	}
	return t || f
}

func main() {
	fmt.Println(check_if_string_can_break_another_string("szy", "cid"))
}
