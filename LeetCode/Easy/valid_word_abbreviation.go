/*
A string can be abbreviated by replacing any number of non-adjacent, non-empty substrings with their lengths.
The lengths should not have leading zeros.

For example, a string such as "substitution" could be abbreviated as (but not limited to):

"s10n" ("s ubstitutio n")
"sub4u4" ("sub stit u tion")
"12" ("substitution")
"su3i1u2on" ("su bst i t u ti on")
"substitution" (no substrings replaced)
The following are not valid abbreviations:

"s55n" ("s ubsti tutio n", the replaced substrings are adjacent)
"s010n" (has leading zeros)
"s0ubstitution" (replaces an empty substring)
Given a string word and an abbreviation abbr, return whether the string matches the given abbreviation.

A substring is a contiguous non-empty sequence of characters within a string.
*/

package main

import (
	"fmt"
	"strconv"
)

func is_valid(word, abbr string) bool {
	m, n := len(word), len(abbr)
	i, k := 0, 0
	val := ""
	if n > m {
		return false
	}
	for i = 0; i < n && k < m; {
		if abbr[i] >= '0' && abbr[i] <= '9' {
			for i < n && abbr[i] >= '0' && abbr[i] <= '9' {
				val += string(abbr[i])
				i++
			}
			if val[0] == '0' {
				return false
			}
			num, _ := strconv.Atoi(val)
			val = ""
			k += num
			i--
		} else {
			if word[k] != abbr[i] {
				return false
			}
			k++
		}
		i++
	}

	if k > m || k <= m-1 || i < n {
		return false
	}
	return true
}

func main() {
	fmt.Println(is_valid("internationalization", "i12iz4n"))
	fmt.Println(is_valid("apple", "a2e"))
	fmt.Println(is_valid("hi", "2i"))
	fmt.Println(is_valid("internationalization", "i5a11o1"))
	fmt.Println(is_valid("hi", "hi1"))
	fmt.Println(is_valid("a", "1"))
}
