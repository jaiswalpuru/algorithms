/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
*/

package main

func valid_sudoku(arr [][]byte) bool {
	for i := 0; i < 9; i++ {
		mp := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if arr[i][j] >= '1' && arr[i][j] <= '9' {
				if _, ok := mp[arr[i][j]]; ok {
					return false
				}
				mp[arr[i][j]] = struct{}{}
			}
		}

		mp = make(map[byte]struct{})
		for k := 0; k < 9; k++ {
			if arr[k][i] >= '1' && arr[k][i] <= '9' {
				if _, ok := mp[arr[k][i]]; ok {
					return false
				}
				mp[arr[k][i]] = struct{}{}
			}
		}
	}

	for i := 0; i < 6; i += 3 {
		mp := make(map[byte]struct{})
		for j := i; j < i+3; j++ {
			for k := i; k < i+3; k++ {
				if arr[j][k] >= '1' && arr[j][k] <= '9' {
					if _, ok := mp[arr[j][k]]; ok {
						return false
					}
					mp[arr[j][k]] = struct{}{}
				}
			}
		}
	}
	return true
}

func main() {
}
