package main

import "fmt"

//1->a, 2->e, 3->i, 4->o, 5->u
//-----------------------backtrack-----------------------
func count_vowels(n int) int {
	return _count_vowels(n, 1)
}

func _count_vowels(n int, vowels int) int {
	if n == 0 {
		return 1
	}
	res := 0
	for i := vowels; i <= 5; i++ {
		res += _count_vowels(n-1, i)
	}
	return res
}

//---------------------------------------------------------

//-----------------------Recursive approach again--------------------- (refer to LC soln)

func count_vowels_recursion(n int) int {
	return _count_vowels_recursion(n, 5)
}

func _count_vowels_recursion(n int, vowels int) int {
	if n == 1 {
		return vowels
	}
	if vowels == 1 {
		return 1
	}
	return _count_vowels_recursion(n-1, vowels) + _count_vowels_recursion(n, vowels-1)
}

//---------------------------------------------------------------------

//-----------------------------using top down memoization--------------
func count_vowels_memoization(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6)
	}
	return _memo(n, 5, &dp)
}

func _memo(n int, vowels int, dp *[][]int) int {
	if n == 1 {
		return vowels
	}
	if vowels == 1 {
		return 1
	}

	if (*dp)[n][vowels] != 0 {
		return (*dp)[n][vowels]
	}
	res := _memo(n-1, vowels, dp) + _memo(n, vowels-1, dp)
	(*dp)[n][vowels] = res
	return res
}

//---------------------------------------------------------------------

//------------------------------bottom up approach---------------------
func count_vowels_bottom_up(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6)
	}

	for i := 1; i <= 5; i++ {
		dp[1][i] = i
	}

	for i := 2; i <= n; i++ {
		dp[i][1] = 1
		for vow := 2; vow <= 5; vow++ {
			dp[i][vow] = dp[i-1][vow] + dp[i][vow-1]
		}
	}

	return dp[n][5]
}

//---------------------------------------------------------------------

func main() {
	fmt.Println(count_vowels_memoization(33))
}
