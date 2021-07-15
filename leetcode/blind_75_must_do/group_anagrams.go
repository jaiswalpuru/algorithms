/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func group_anagrams(arr []string) [][]string {
	mp := make(map[string][]string)
	n := len(arr)
	res := make([][]string, 0)

	for i := 0; i < n; i++ {
		temp := strings.Split(arr[i], "")
		sort.Strings(temp)
		str := strings.Join(temp, "")
		if _, ok := mp[str]; ok {
			mp[str] = append(mp[str], arr[i])
		} else {
			mp[str] = append(mp[str], arr[i])
		}
	}

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}

func main() {
	fmt.Println(group_anagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(group_anagrams([]string{}))
	fmt.Println(group_anagrams([]string{"a"}))
}
