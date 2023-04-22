package main

import (
	"fmt"
	"reflect"
)

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if pLen > sLen {
		return nil
	}

	pCharCnt := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pCharCnt[p[i]]++
	}

	index := []int{}
	sCharCnt := make(map[byte]int)
	for i := 0; i < sLen; i++ {
		sCharCnt[s[i]]++
		if i >= pLen {
			sCharCnt[s[i-pLen]]--
			if sCharCnt[s[i-pLen]] == 0 {
				delete(sCharCnt, s[i-pLen])
			}
		}
		if len(sCharCnt) == len(pCharCnt) {
			if reflect.DeepEqual(sCharCnt, pCharCnt) {
				index = append(index, i-pLen+1)
			}
		}
	}
	return index
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
