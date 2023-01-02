package main

import "fmt"

//-------------brute force------------------
func longest_palindromic_subsequence_two(s string) int {
	return recur(s, 0, len(s)-1, '_')
}

func recur(s string, l, r int, prev byte) int {
	if l >= r {
		return 0
	}
	if s[l] == prev {
		return recur(s, l+1, r, prev)
	}
	if s[r] == prev {
		return recur(s, l, r-1, prev)
	}
	if s[l] == s[r] {
		return 2 + recur(s, l+1, r-1, s[l])
	} else {
		return max(recur(s, l+1, r, prev), recur(s, l, r-1, prev))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------brute force------------------

//-------------memo------------------
func longest_palindromic_subsequence_two_memo(s string) int {
	memo := make([][][27]int, len(s))
	for i := 0; i < len(s); i++ {
		memo[i] = make([][27]int, len(s))
		for j := 0; j < len(s); j++ {
			for k := 0; k < 27; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	return _recur(s, 0, len(s)-1, '{', &memo)
}

func _recur(s string, l, r int, prev byte, memo *[][][27]int) int {
	if l >= r {
		return 0
	}
	char := int(prev) - int('a')
	if (*memo)[l][r][char] != -1 {
		return (*memo)[l][r][char]
	}
	if s[l] == prev {
		(*memo)[l][r][char] = _recur(s, l+1, r, prev, memo)
		return (*memo)[l][r][char]
	}
	if s[r] == prev {
		(*memo)[l][r][char] = _recur(s, l, r-1, prev, memo)
		return (*memo)[l][r][char]
	}
	if s[l] == s[r] {
		(*memo)[l][r][char] = 2 + _recur(s, l+1, r-1, s[l], memo)
	} else {
		(*memo)[l][r][char] = max(_recur(s, l+1, r, prev, memo), _recur(s, l, r-1, prev, memo))
	}
	return (*memo)[l][r][char]
}

//-------------memo------------------

func main() {
	fmt.Println(longest_palindromic_subsequence_two_memo("bbabab"))
}
