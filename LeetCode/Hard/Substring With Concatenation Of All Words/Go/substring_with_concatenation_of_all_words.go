package main

import "fmt"

func substring_with_concatenation_of_all_words(s string, words []string) []int {
	n := len(s)
	k := len(words)
	word_len := len(words[0])
	sub_str_size := k * word_len
	hash_map := make(map[string]int)
	for i := 0; i < len(words); i++ {
		hash_map[words[i]]++
	}
	res := []int{}
	for i := 0; i < n-sub_str_size+1; i++ {
		if is_valid(i, s, word_len, sub_str_size, hash_map, k) {
			res = append(res, i)
		}
	}
	return res
}

func is_valid(start int, s string, word_len, sub_str_size int, hash_map map[string]int, k int) bool {
	mp := make(map[string]int)
	for k, v := range hash_map {
		mp[k] = v
	}
	word_used := 0
	for i := start; i < start+sub_str_size; i += word_len {
		sub_str := s[i : i+word_len]
		if mp[sub_str] != 0 {
			mp[sub_str]--
			word_used++
		} else {
			break
		}
	}
	return word_used == k
}

func main() {
	fmt.Println(substring_with_concatenation_of_all_words("barfoothefoobarman", []string{"foo", "bar"}))
}
