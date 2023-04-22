package main

import (
	"fmt"
	"strings"
)

func are_close(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n != m {
		return false
	}

	str1 := strings.Split(s1, "")
	str2 := strings.Split(s2, "")

	mp_str1 := make(map[string]int)
	mp_str2 := make(map[string]int)
	for i := 0; i < n; i++ {
		mp_str1[str1[i]]++
		mp_str2[str2[i]]++
	}

	for k := range mp_str1 {
		if _, ok := mp_str2[k]; !ok {
			return false
		}
	}

	visited_1 := make(map[int]int)
	visited_2 := make(map[int]int)
	for _, v := range mp_str1 {
		visited_1[v]++
	}

	for _, v := range mp_str2 {
		visited_2[v]++
	}

	for k := range visited_2 {
		if visited_1[k] != visited_2[k] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(are_close("cabbba", "aabbss"))
}
