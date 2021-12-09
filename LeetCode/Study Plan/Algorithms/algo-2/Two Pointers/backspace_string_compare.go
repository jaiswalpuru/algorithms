/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.
*/
package main

import "fmt"

func backspace_string_compare(s, t string) bool {
	st, tt := []byte{}, []byte{}
	n, m := len(s), len(t)
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, s[i])
		}
	}

	for i := 0; i < m; i++ {
		if t[i] == '#' {
			if len(tt) > 0 {
				tt = tt[:len(tt)-1]
			}
		} else {
			tt = append(tt, t[i])
		}
	}
	if string(st) == string(tt) {
		return true
	}
	return false
}

func main() {
	fmt.Println(backspace_string_compare("ab#c", "ad#c"))
	fmt.Println(backspace_string_compare("ab##", "c#d#"))
	fmt.Println(backspace_string_compare("a##c", "#a#c"))
	fmt.Println(backspace_string_compare("a#c", "b"))
}
