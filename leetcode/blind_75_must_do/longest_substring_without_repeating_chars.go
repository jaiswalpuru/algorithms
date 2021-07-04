/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.


Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.


Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Example 4:

Input: s = ""
Output: 0
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func len_of_longest_increasing_subseq(s string) int {
	mp := make(map[string]int)
	n := len(s)
	length := 0
	i, j := 0, 0
	for i < n {
		if _, ok := mp[string(s[i])]; ok {
			t := i - j
			length = max(length, t)
			if j <= (mp[string(s[i])]) {
				j = mp[string(s[i])] + 1
			}
		}
		mp[string(s[i])] = i
		i++
	}
	return max(length, i-j)
}

func main() {
	fmt.Println(len_of_longest_increasing_subseq(""))
	fmt.Println(len_of_longest_increasing_subseq("abcabcbb"))
	fmt.Println(len_of_longest_increasing_subseq("bbbbb"))
	fmt.Println(len_of_longest_increasing_subseq("pwwkew"))
	fmt.Println(len_of_longest_increasing_subseq("aab"))
	fmt.Println(len_of_longest_increasing_subseq("dvdf"))
	fmt.Println(len_of_longest_increasing_subseq("cdd"))
	fmt.Println(len_of_longest_increasing_subseq("abba"))
	fmt.Println(len_of_longest_increasing_subseq("ckilbkd"))
}
