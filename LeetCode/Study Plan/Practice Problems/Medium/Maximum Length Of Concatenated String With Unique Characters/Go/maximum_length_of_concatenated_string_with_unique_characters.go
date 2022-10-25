package main

import (
	"fmt"
	"strings"
)

func check_dup(str string) bool {
	n := len(str)
	mp := make(map[byte]int)
	for i := 0; i < n; i++ {
		mp[str[i]]++
		if mp[str[i]] > 1 {
			return true
		}
	}
	return false
}

func maximum_length_of_concatenated_string_with_unique_characters(arr []string) int {
	res := 0
	backtrack(arr, &res, 0, []string{})
	return res
}

func backtrack(arr []string, res *int, ind int, str []string) {
	n := len(arr)
	cnt := 0
	for i := ind; i < n; i++ {
		str = append(str, arr[i])
		val := strings.Join(str, "")
		if !check_dup(val) {
			cnt = len(val)
			backtrack(arr, res, i+1, str)
		}
		*res = max(*res, cnt)
		str = str[:len(str)-1]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximum_length_of_concatenated_string_with_unique_characters([]string{"un", "iq", "ue"}))
}
