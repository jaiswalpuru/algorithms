/*
The edit distance between two strings refers to the minimum number of character insertions,
deletions, and substitutions required to change one string to the other. For example, the edit
distance between “kitten” and “sitting” is three: substitute the “k” for “s”, substitute the “e” for “i”,
and append a “g”.

Given two strings, compute the edit distance between them.
*/

package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func computeDistance(s1, s2 string) int {
	distance := 0
	n, m := len(s1), len(s2)
	if n != m {
		distance += abs(n - m)
	}
	if n < m {
		n, m = m, n
		s1, s2 = s2, s1
	}
	mp := make(map[rune]int)
	for i := 0; i < n; i++ {
		mp[rune(s1[i])]++
	}
	for i := 0; i < m; i++ {
		if _, ok := mp[rune(s2[i])]; ok {
			mp[rune(s2[i])]--
		} else {
			distance++
		}
	}
	return distance
}

func main() {
	s1, s2 := "kitten", "sitting"
	fmt.Println(computeDistance(s1, s2))
}
