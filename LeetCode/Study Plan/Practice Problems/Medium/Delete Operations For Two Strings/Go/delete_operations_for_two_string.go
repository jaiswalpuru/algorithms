package main

import "fmt"

//-----------------------brute force recursion-------------------------
func delete_operations_for_two_string(s1, s2 string) int {
	return _recur(0, 0, s1, s2)
}

func _recur(i, j int, s1, s2 string) int {
	if i == len(s1) {
		return len(s2) - j
	}

	if j == len(s2) {
		return len(s1) - i
	}

	if s1[i] == s2[j] {
		return _recur(i+1, j+1, s1, s2)
	} else {
		del_s1 := 1 + _recur(i+1, j, s1, s2)
		del_s2 := 1 + _recur(i, j+1, s1, s2)
		return min(del_s1, del_s2)
	}
}

//-----------------------brute force recursion-------------------------

//-----------------------memoizations-------------------------
func delete_operations_for_two_string_eff(s1, s2 string) int {
	n, m := len(s1), len(s2)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}
	return _memo(0, 0, s1, s2, &memo)
}

func _memo(i, j int, s1, s2 string, memo *[][]int) int {
	if i == len(s1) {
		return len(s2) - j
	}

	if j == len(s2) {
		return len(s1) - i
	}

	if (*memo)[i][j] != -1 {
		return (*memo)[i][j]
	}

	if s1[i] == s2[j] {
		(*memo)[i][j] = _memo(i+1, j+1, s1, s2, memo)
	} else {
		del_s1, del_s2 := 1+_memo(i+1, j, s1, s2, memo), 1+_memo(i, j+1, s1, s2, memo)
		(*memo)[i][j] = min(del_s1, del_s2)
	}
	return (*memo)[i][j]
}

//-----------------------memoizations-------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(delete_operations_for_two_string_eff("sea", "eat"))
}
