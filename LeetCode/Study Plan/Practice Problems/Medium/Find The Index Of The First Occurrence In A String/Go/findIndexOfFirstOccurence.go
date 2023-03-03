package main

import "fmt"

func strStr(haystack string, needle string) int {
	hSize := len(haystack)
	nSize := len(needle)
	i := 0

	for i < hSize {
		start := i
		k := 0
		for start < hSize && k < nSize && haystack[start] == needle[k] {
			start++
			k++
		}

		if k == nSize {
			return i
		}
		i++
	}

	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}
