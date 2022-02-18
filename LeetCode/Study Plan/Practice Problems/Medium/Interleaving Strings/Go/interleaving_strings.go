package main

import "fmt"

//--------------------------interleaving strings--------------------------(TLE)
func interleaving_string(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return _interleaving_string(s1, s2, s3)
}

func _interleaving_string(s1, s2, s3 string) bool {
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}

	if s1[0] == s2[0] && s1[0] == s3[0] {
		return _interleaving_string(s1[1:], s2, s3[1:]) || _interleaving_string(s1, s2[1:], s3[1:])
	} else if s2[0] == s3[0] {
		return _interleaving_string(s1, s2[1:], s3[1:])
	} else if s1[0] == s3[0] {
		return _interleaving_string(s1[1:], s2, s3[1:])
	}
	return false
}

//------------------------------------------------------------------------

//----------------------------------Memoization-----------------------------------
func interleaving_string_memo(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	n, m := len(s1), len(s2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}

	return _interleaving_string_memo(s1, s2, s3, 0, 0, &dp)
}

func _interleaving_string_memo(s1, s2, s3 string, i, j int, dp *[][]int) bool {
	if i == len(s1) {
		return s2[j:] == s3
	}

	if j == len(s2) {
		return s1[i:] == s3
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j] == 1
	}

	if s1[i] == s2[j] && s1[i] == s3[0] {
		val := _interleaving_string_memo(s1, s2, s3[1:], i+1, j, dp) || _interleaving_string_memo(s1, s2, s3[1:], i, j+1, dp)
		if val {
			(*dp)[i][j] = 1
		} else {
			(*dp)[i][j] = 0
		}
	} else if s1[i] == s3[0] {
		val := _interleaving_string_memo(s1, s2, s3[1:], i+1, j, dp)
		if val {
			(*dp)[i][j] = 1
		} else {
			(*dp)[i][j] = 0
		}
	} else if s2[j] == s3[0] {
		val := _interleaving_string_memo(s1, s2, s3[1:], i, j+1, dp)
		if val {
			(*dp)[i][j] = 1
		} else {
			(*dp)[i][j] = 0
		}
	} else {
		(*dp)[i][j] = 0
	}
	return (*dp)[i][j] == 1
}

//-------------------------------------------------------------------------------------

//--------------------------------------dp-----------------------------------------------
func interleaving_strings(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	n, m := len(s1), len(s2)
	dp := make([][]bool, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[n][m]
}

//--------------------------------------dp-----------------------------------------------

func main() {
	fmt.Println(interleaving_strings("aabcc", "dbbca", "aadbbcbcac"))
}
