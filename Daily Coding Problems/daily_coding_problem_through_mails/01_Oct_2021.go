/*
Given a start word, an end word, and a dictionary of valid words, find the shortest transformation
sequence from start to end such that only one letter is changed at each step of the sequence, and
each transformed word exists in the dictionary. If there is no possible transformation, return null.
Each word in the dictionary have the same length as start and end and is lowercase.

For example, given start = "dog", end = "cat", and dictionary = {"dot", "dop", "dat", "cat"},
return ["dog", "dot", "dat", "cat"].

Given start = "dog", end = "cat", and dictionary = {"dot", "tod", "dat", "dar"}, return null as there is
no possible transformation from dog to cat.
*/

package main

import "fmt"

func shorttest_transformation(begin, end string, words []string) []string {
	n := len(words)
	if n == 0 {
		return nil
	}

	mp := make(map[string]struct{})
	for i := 0; i < n; i++ {
		mp[words[i]] = struct{}{}
	}

	//If the word is not present in the array of words
	if _, ok := mp[end]; !ok {
		return nil
	}

	m := len(begin)
	for i := 0; i < m; i++ {
		temp := begin
		for j := 0; j < 26; j++ {

		}
	}

}

func main() {
	fmt.Println(shorttest_transformation("dog", "cat", []string{"dot", "dop", "dat", "cat"}))
}
