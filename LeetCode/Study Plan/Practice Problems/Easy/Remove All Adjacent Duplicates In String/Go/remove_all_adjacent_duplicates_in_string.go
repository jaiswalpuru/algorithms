package main

import "fmt"

func remove_all_adjacent_duplicates_in_string(s string) string {
	st := []byte(s)
	stack := []byte{}
	n := len(st)
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, st[i])
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == st[i] {
			for len(stack) > 0 && stack[len(stack)-1] == st[i] {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, st[i])
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(remove_all_adjacent_duplicates_in_string("abbaca"))
}
