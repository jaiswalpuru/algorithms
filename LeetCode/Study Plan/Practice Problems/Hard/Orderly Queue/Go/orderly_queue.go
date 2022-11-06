package main

import (
	"fmt"
	"sort"
)

func orderly_queue(s string, k int) string {
	n := len(s)
	if k == 1 {
		ans := s
		for i := 0; i < n; i++ {
			temp := s[i:] + s[0:i]
			if temp < ans {
				ans = temp
			}
		}
		return ans
	} else {
		str := []byte(s)
		sort.Slice(str, func(i, j int) bool { return str[i] < str[j] })
		return string(str)
	}
}

func main() {
	fmt.Println(orderly_queue("cba", 1))
}
