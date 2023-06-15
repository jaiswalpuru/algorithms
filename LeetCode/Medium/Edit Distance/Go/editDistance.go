package main

import "fmt"

//------------------------brute force try to generate all possibilities---------------------
func minDistanceBruteForce(word1 string, word2 string) int {
	return editDistanceBruteForce(0, 0, word1, word2)
}

func editDistanceBruteForce(i, j int, word1, word2 string) int {
	if i == len(word1) {
		return len(word2[j:])
	}

	if j == len(word2) {
		return len(word1[i:])
	}

	if word1[i] != word2[j] {
		insert := 1 + editDistanceBruteForce(i, j+1, word1, word2)
		remove := 1 + editDistanceBruteForce(i+1, j, word1, word2)
		replace := 1 + editDistanceBruteForce(i+1, j+1, word1, word2)
		return min(insert, min(remove, replace))
	} else {
		return editDistanceBruteForce(i+1, j+1, word1, word2)
	}
}

//------------------------brute force try to generate all possibilities---------------------

//--------------------------------memoization----------------------------
func minDistance(word1, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}

	return editDistance(0, 0, word1, word2, &dp)
}

func editDistance(i, j int, word1, word2 string, dp *[][]int) int {
	if i == len(word1) {
		return len(word2) - j
	}

	if j == len(word2) {
		return len(word1) - i
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	if word1[i] != word2[j] {
		insert := 1 + editDistance(i, j+1, word1, word2, dp)
		remove := 1 + editDistance(i+1, j, word1, word2, dp)
		replace := 1 + editDistance(i+1, j+1, word1, word2, dp)
		(*dp)[i][j] = min(insert, min(remove, replace))
	} else {
		(*dp)[i][j] = editDistance(i+1, j+1, word1, word2, dp)
	}
	return (*dp)[i][j]
}

//--------------------------------memoization----------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
