/*
	Given a 2D matrix of characters and a target word, write a function that returns whether the word can be
	found in the matrix by going left-to-right, or up-to-down.

	For example, given the following matrix:

	[['F', 'A', 'C', 'I'],
 	['O', 'B', 'Q', 'P'],
 	['A', 'N', 'O', 'B'],
 	['M', 'A', 'S', 'S']]
	and the target word 'FOAM', you should return true, since it's the leftmost column. Similarly,
	given the target word 'MASS', you should return true, since it's the last row.

	Approach :
		If cell has first character, then we one by one try all possible 8 directions from that cell to match.

*/

package main

import "fmt"

var (
	x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y = []int{-1, 0, 1, -1, 1, -1, 0, 1}
)

func pattern_search(arr [][]rune, word []rune) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if search(arr, i, j, word, len(arr), len(arr[i])) {
				fmt.Println("Pattern found at : ", i, j)
			}
		}
	}
}

func search(arr [][]rune, row, col int, word []rune, r, c int) bool {
	//if the first character doesn't match
	if arr[row][col] != word[0] {
		return false
	}

	l := len(word)
	//check in all possible 8 directions
	for i := 0; i < 8; i++ {
		var k int
		rd, cd := row+x[i], col+y[i]
		for k = 1; k < l; k++ {
			if rd >= r || rd < 0 || cd >= c || cd < 0 {
				break
			}
			if arr[rd][cd] != word[k] {
				break
			}
			rd += x[i]
			cd += y[i]
		}
		if k == l {
			return true
		}
	}
	return false
}

func main() {
	arr := [][]rune{
		{'F', 'A', 'C', 'I'},
		{'O', 'B', 'Q', 'P'},
		{'A', 'N', 'O', 'B'},
		{'M', 'A', 'S', 'S'},
	}
	word := "FOAM"
	word_rune := []rune(word)

	pattern_search(arr, word_rune)
}
