package main

import (
	"fmt"
	"sort"
)

func group_anagrams(strs []string) [][]string {
	n := len(strs)
	mp := make(map[string][]string)
	for i := 0; i < n; i++ {
		word := []byte(strs[i])
		sort.Slice(word, func(i, j int) bool { return word[i] < word[j] })
		mp[string(word)] = append(mp[string(word)], strs[i])
	}
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(group_anagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
