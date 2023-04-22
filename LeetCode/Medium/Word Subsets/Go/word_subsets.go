package main

import (
	"fmt"
)

func word_subsets(word1 []string, word2 []string) []string {
	res := []string{}
	cnt := count("")
	n, m := len(word1), len(word2)
	for i := 0; i < m; i++ {
		word_cnt := count(word2[i])
		for i := 0; i < 26; i++ {
			cnt[i] = max(cnt[i], word_cnt[i])
		}
	}
	for i := 0; i < n; i++ {
		word_cnt := count(word1[i])
		f := true
		for j := 0; j < 26; j++ {
			if word_cnt[j] < cnt[j] {
				f = false
				break
			}
		}
		if f {
			res = append(res, word1[i])
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

func count(str string) []int {
	cnt := make([]int, 26)
	for i := 0; i < len(str); i++ {
		cnt[str[i]-'a']++
	}
	return cnt
}

func main() {
	fmt.Println(word_subsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
}
