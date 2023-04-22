/*
Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal,
otherwise, return false.

Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters
at s[i] and s[j].

For example, swapping at indices 0 and 2 in "abcd" results in "cbad".
*/

package main

import "fmt"

func buddy_strings(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n != m {
		return false
	}

	mp_s1 := make(map[byte]int)

	for i := 0; i < n; i++ {
		mp_s1[s1[i]]++
	}

	if s1 == s2 {
		for _, v := range mp_s1 {
			if v > 1 {
				return true
			}
		}
		return false
	}

	c1, c2 := -1, -1

	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			if c1 == -1 {
				c1 = i
			} else if c2 == -1 {
				c2 = i
			} else {
				return false
			}
		}
	}

	return c2 != -1 && s1[c1] == s2[c2] && s1[c2] == s2[c1]
}

func main() {
	fmt.Println(buddy_strings("ab", "ba"))
	fmt.Println(buddy_strings("ab", "ab"))
	fmt.Println(buddy_strings("aa", "aa"))
}
