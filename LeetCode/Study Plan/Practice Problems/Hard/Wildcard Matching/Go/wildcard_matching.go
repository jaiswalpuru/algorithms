package main

import "fmt"

//----------------------recursion-----------------------
func wildcard_matching(s, p string) bool {
	return _recur(s, p, 0, 0)
}

func _recur(s, p string, i, j int) bool {
	if j == len(p) {
		return i == len(s)
	}
	match := false
	if i == len(s) {
		if p[j] != '*' {
			return false
		}
		match = _recur(s, p, i, j+1)
	} else if p[j] == '?' {
		match = _recur(s, p, i+1, j+1)
	} else if p[j] == '*' {
		match = _recur(s, p, i+1, j) || _recur(s, p, i, j+1) || _recur(s, p, i+1, j+1)
	} else {
		if s[i] == p[j] {
			match = _recur(s, p, i+1, j+1)
		}
	}
	return match
}

//----------------------recursion-----------------------

//----------------------memo-----------------------
func wildcard_matching_memo(s, p string) bool {
	memo := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		memo[i] = make([]int, len(p)+1)
		for j := 0; j < len(p)+1; j++ {
			memo[i][j] = -1
		}
	}
	return recur(s, p, 0, 0, &memo)
}

func recur(s, p string, i, j int, memo *[][]int) bool {
	if (*memo)[i][j] != -1 {
		return (*memo)[i][j] == 1
	}

	if j == len(p) {
		return i == len(s)
	}
	match := false
	if i == len(s) {
		if p[j] != '*' {
			(*memo)[i][j] = 0
			return false
		}
		match = recur(s, p, i, j+1, memo)
	} else if p[j] == '?' {
		match = recur(s, p, i+1, j+1, memo)
	} else if p[j] == '*' {
		match = match || recur(s, p, i+1, j, memo) || recur(s, p, i, j+1, memo) || recur(s, p, i+1, j+1, memo)
	} else {
		if s[i] == p[j] {
			match = recur(s, p, i+1, j+1, memo)
		}
	}
	if match {
		(*memo)[i][j] = 1
	} else {
		(*memo)[i][j] = 0
	}
	return match
}

//----------------------memo-----------------------

func main() {
	fmt.Println(wildcard_matching_memo("aa", "a"))
}
