package main

import "fmt"

func maximum_word_prod_length(arr []string) int {
	n := len(arr)

	res := 0
	for i := 0; i < n; i++ {
		mp := make(map[string]struct{})
		m := len(arr[i])
		for j := 0; j < m; j++ {
			mp[string(arr[i][j])] = struct{}{}
		}

		f := false
		var t int
		for j := i + 1; j < n; j++ {
			t = len(arr[j])
			for k := 0; k < t; k++ {
				if _, ok := mp[string(arr[j][k])]; ok {
					f = true
					break
				}
			}
			if f == false {
				res = max(res, m*t)
			}
			f = false
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximum_word_prod_length([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}
