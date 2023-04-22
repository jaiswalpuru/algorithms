package main

import "fmt"

func count_common_words_with_once_occurrence(words1, words2 []string) int {
	word1_cnt := make(map[string]int)
	word2_cnt := make(map[string]int)
	cnt := 0
	for i := 0; i < len(words1); i++ {
		word1_cnt[words1[i]]++
	}
	for i := 0; i < len(words2); i++ {
		word2_cnt[words2[i]]++
	}
	for k, v := range word1_cnt {
		if v == 1 && word2_cnt[k] == 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(count_common_words_with_once_occurrence([]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}))
}
