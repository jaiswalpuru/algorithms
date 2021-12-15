/*
Write an algorithm to justify text. Given a sequence of words and an integer line length k,
return a list of strings which represents each line, fully justified.

More specifically, you should have as many words as possible in each line.
There should be at least one space between each word.
Pad extra spaces when necessary so that each line has exactly length k.
Spaces should be distributed as equally as possible, with the extra spaces, if any,
distributed starting from the left.

If you can only fit one word on a line, then you should pad the right-hand side with spaces.

Each word is guaranteed not to be longer than k.

For example, given the list of words ["the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"]
and k = 16, you should return the following:

["the  quick brown", # 1 extra space on the left
"fox  jumps  over", # 2 extra spaces distributed evenly
"the   lazy   dog"] # 4 extra spaces distributed evenly
*/

package main

import (
	"fmt"
	"math"
	"strings"
)

func getMinLine(l []string, word string) string {
	s := strings.Join(l, " ")
	s += " " + word
	return s
}

func groupWords(words []string, k int) [][]string {
	groups := make([][]string, 0)
	currLine := []string{}

	for _, word := range words {
		if len(getMinLine(currLine, word)) > k {
			groups = append(groups, currLine)
			currLine = []string{}
		}
		currLine = append(currLine, word)
	}
	groups = append(groups, currLine)
	return groups
}

func sumOfWordLen(words []string) int {
	sum := 0
	for i := 0; i < len(words); i++ {
		sum += len(words[i])
	}
	return sum
}

func justify(words []string, k int) string {
	word := ""
	if len(words) == 1 {
		word = words[0]
		spaces := k - len(word)
		for i := 0; i < spaces; i++ {
			word = word + " "
		}
		return word
	}
	spaceToDistribute := k - sumOfWordLen(words)
	numOfSpaces := len(words) - 1
	smallestSpace := int(math.Floor(float64(spaceToDistribute) / float64(numOfSpaces)))
	leftOverSpace := spaceToDistribute - (numOfSpaces * smallestSpace)

	justifiedWords := []string{}

	for _, word := range words {
		justifiedWords = append(justifiedWords, word)
		currSpace := ""
		for i := 0; i < smallestSpace; i++ {
			currSpace += " "
		}
		if leftOverSpace > 0 {
			currSpace += " "
			leftOverSpace--
		}
		justifiedWords = append(justifiedWords, currSpace)
	}
	return strings.Join(justifiedWords, "")
}

func main() {

	words := []string{
		"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog",
	}
	k := 16
	g := groupWords(words, k)
	for i := 0; i < len(g); i++ {
		fmt.Println(justify(g[i], k))
	}
}
