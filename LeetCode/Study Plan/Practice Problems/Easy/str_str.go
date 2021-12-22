/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and
Java's indexOf().
*/

package main

import "fmt"

func strstr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	ind := -1

	if needle == "" && haystack == "" {
		return 0
	}

	if haystack == "" {
		return ind
	}

	if needle == "" {
		return 0
	}

	for i := 0; i < n; i++ {
		if haystack[i] == needle[0] {
			if i+m <= n && haystack[i:i+m] == needle {
				ind = i
				break
			}
		}
	}

	return ind
}

func main() {
	fmt.Println(strstr("hello", "ll"))
	fmt.Println(strstr("aaaaa", "bba"))
	fmt.Println(strstr("", ""))
	fmt.Println(strstr("", "aa"))
	fmt.Println(strstr("a", "a"))
}
