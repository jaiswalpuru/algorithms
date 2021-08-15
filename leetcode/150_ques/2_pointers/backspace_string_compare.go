/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors.
'#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

Example 1:
Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:
Input: s = "a##c", t = "#a#c"
Output: true
Explanation: Both s and t become "c".

Example 4:
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
*/

package main

import "fmt"

func backspace_string_compare(s, t string) bool {
	n, m := len(s), len(t)
	s1 := []byte{}
	s2 := []byte{}

	for i := 0; i < n; i++ {
		if s[i] != '#' {
			s1 = append(s1, s[i])
		} else {
			if len(s1) != 0 {
				s1 = s1[:len(s1)-1]
			}
		}
	}

	for i := 0; i < m; i++ {
		if t[i] != '#' {
			s2 = append(s2, t[i])
		} else {
			if len(s2) != 0 {
				s2 = s2[:len(s2)-1]
			}
		}
	}

	if len(s1) != len(s2) {
		return false
	} else {
		k := len(s1)
		for i := 0; i < k; i++ {
			if s1[i] != s2[i] {
				return false
			}
		}

	}

	return true
}

func main() {
	fmt.Println(backspace_string_compare("ab#c", "ad#c"))
	fmt.Println(backspace_string_compare("ab##", "c#d#"))
	fmt.Println(backspace_string_compare("a#c", "b"))
	fmt.Println(backspace_string_compare("abc#", "bac#"))
}
