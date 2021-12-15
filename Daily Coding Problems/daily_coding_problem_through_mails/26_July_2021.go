/*
Given a string and a set of characters, return the shortest substring containing all the characters in the set.

For example, given the string "figehaeci" and the set of characters {a, e, i}, you should return "aeci".

If there is no substring containing all the characters in the set, return null.
*/

package main

import (
	"fmt"
	"strings"
)

func get_count(str string) map[byte]int {
	mp := make(map[byte]int)
	n := len(str)
	for i := 0; i < n; i++ {
		mp[str[i]]++
	}
	return mp
}

func shortest_substring(str string, char string) string {
	n, m := len(str), len(char)
	if n == 0 || m == 0 {
		return ""
	}
	mp_char := get_count(char)
	mp_str := make(map[byte]int)
	formed, reqd := 0, len(mp_char)
	ans := []int{-1, 0, 0} //length, start index, end index

	l, r := 0, 0

	for r < n {
		mp_str[str[r]]++
		if mp_str[str[r]] == mp_char[str[r]] {
			formed++
		}

		for l <= r && formed == reqd {
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}
			mp_str[str[l]]--
			if mp_str[str[l]] < mp_char[str[l]] {
				formed--
			}
			l++
		}
		r++
	}

	if ans[0] == -1 {
		return ""
	}
	return str[ans[1] : ans[2]+1]
}

func main() {
	fmt.Println(shortest_substring("figehaeci", strings.Join([]string{"a", "e", "i"}, "")))
}
