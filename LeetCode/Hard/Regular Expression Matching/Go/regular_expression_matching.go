package main

import "fmt"

//-----------------------recursion------------------------
func regular_expression_matching(s, p string) bool {
	return recur(s, p, 0, 0)
}

func recur(s, p string, i, j int) bool {
	ans := false
	if j == len(p) {
		ans = i == len(s)
	} else {
		first_match := false
		if i < len(s) {
			first_match = p[j] == '.' || s[i] == p[j]
		}
		if j+1 < len(p) && p[j+1] == '*' {
			ans = recur(s, p, i, j+2) || first_match && recur(s, p, i+1, j)
		} else {
			ans = first_match && recur(s, p, i+1, j+1)
		}
	}
	return ans
}

//-----------------------recursion------------------------

func regular_expression_matching_memo(s, p string) bool {
	memo := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		memo[i] = make([]int, len(p)+1)
		for j := 0; j < len(p)+1; j++ {
			memo[i][j] = -1
		}
	}
	return _memo(s, p, 0, 0, &memo)
}

func _memo(s, p string, i, j int, memo *[][]int) bool {
	if (*memo)[i][j] != -1 {
		return (*memo)[i][j] == 1
	}
	ans := false
	if j == len(p) {
		ans = i == len(s)
	} else {
		first_match := false
		if i < len(s) {
			first_match = p[j] == '.' || p[j] == s[i]
		}
		if j+1 < len(p) && p[j+1] == '*' {
			ans = _memo(s, p, i, j+2, memo) || first_match && _memo(s, p, i+1, j, memo)
		} else {
			ans = first_match && _memo(s, p, i+1, j+1, memo)
		}
	}
	if ans == false {
		(*memo)[i][j] = 0
	} else {
		(*memo)[i][j] = 1
	}
	return ans
}

func main() {
	fmt.Println(regular_expression_matching_memo("aa", "a"))
}
