package main

import "fmt"

//-------------brute force ------------------
func recursive_multiply(n, m int) int {
	if n > m {
		n, m = m, n
	}
	return min_prod(n, m)
}

func min_prod(small, big int) int {
	if small == 0 {
		return 0
	}
	if small == 1 {
		return big
	}
	s := small >> 2
	s1 := min_prod(s, big)
	s2 := s1
	if small%2 == 1 {
		s2 = min_prod(small-s, big)
	}
	return s1 + s2
}

//-------------brute force ------------------

//-------------memoization------------------
func recursive_multiply_memo(n, m int) int {
	if n > m {
		n, m = m, n
	}
	memo := make([]int, n+1)
	return _memo(n, m, &memo)
}

func _memo(small, big int, memo *[]int) int {
	if small == 0 {
		return 0
	}
	if small == 1 {
		return big
	}
	if (*memo)[small] > 0 {
		return (*memo)[small]
	}
	s := small >> 1
	s1 := _memo(s, big, memo)
	s2 := s1
	if small%2 == 1 {
		s2 = _memo(small-s, big, memo)
	}
	(*memo)[small] = s1 + s2
	return (*memo)[small]
}

//-------------memoization------------------

//-------------better approach------------------
func recursive_multiply_eff(n, m int) int {
	if n > m {
		n, m = m, n
	}
	return _recur(n, m)
}

func _recur(small, big int) int {
	if small == 0 {
		return 0
	}
	if small == 1 {
		return big
	}
	s := small >> 1
	half_prod := _recur(s, big)
	if small%2 == 1 {
		return half_prod + half_prod + big
	}
	return half_prod + half_prod

}

//-------------better approach------------------

func main() {
	fmt.Println(recursive_multiply_eff(3, 9))
}
