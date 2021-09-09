/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words
beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest
transformation sequence from beginWord to endWord, or 0 if no such sequence exists.

Example 1:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

Example 2:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
*/

package main

import "fmt"

var ascii_lower = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
	'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func word_ladder(begin_word, end_word string, word_list []string) int {
	mp := make(map[string]struct{})
	for _, word := range word_list {
		mp[word] = struct{}{}
	}

	if _, ok := mp[end_word]; !ok {
		return 0
	}

	q := []string{begin_word}
	visited := make(map[string]struct{})
	visited[begin_word] = struct{}{}
	length := 1

	for len(q) > 0 {
		q_size := len(q)
		for i := 0; i < q_size; i++ {
			word := q[0]
			q = q[1:]
			for j := 0; j < len(word); j++ {
				for k := 0; k < len(ascii_lower); k++ {
					nxt_word := word[:j] + string(ascii_lower[k]) + word[j+1:]
					if _, ok := mp[nxt_word]; !ok || word == nxt_word {
						continue
					}
					if nxt_word == end_word {
						return length + 1
					}

					if _, ok := visited[nxt_word]; !ok {
						q = append(q, nxt_word)
						visited[nxt_word] = struct{}{}
					}
				}
			}
		}
		length++
	}
	return 0
}

func main() {
	fmt.Println(word_ladder("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(word_ladder("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Println(word_ladder("a", "c", []string{"a", "b", "c"}))
	fmt.Println(word_ladder("hot", "dog", []string{"hot", "dog", "dot"}))
	fmt.Println(word_ladder("ta", "if", []string{"ts", "sc", "ph", "ca", "jr", "hf", "to", "if", "ha", "is", "io", "cf", "ta"}))
}
