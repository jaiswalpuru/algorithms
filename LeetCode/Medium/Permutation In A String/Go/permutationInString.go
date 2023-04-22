package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)

	if s2Len < s1Len {
		return false
	}

	s1CharCnt := make([]int, 26)
	for i := 0; i < s1Len; i++ {
		s1CharCnt[s1[i]-'a']++
	}

	for i := 0; i <= s2Len-s1Len; i++ {
		s2CharCnt := make([]int, 26)
		for j := i; j < i+s1Len; j++ {
			s2CharCnt[s2[j]-'a']++
		}
		if match(s1CharCnt, s2CharCnt) {
			return true
		}
	}

	return false
}

func match(s1CharCnt, s2CharCnt []int) bool {
	for i := 0; i < 26; i++ {
		if s1CharCnt[i] != s2CharCnt[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkInclusion("abc", "cccccbabbbaaaa"))
}
