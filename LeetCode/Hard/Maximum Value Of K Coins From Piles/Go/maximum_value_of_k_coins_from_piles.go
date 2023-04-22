package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//-----------------------------recursion (TLE)-----------------------------
func max_val_of_k_coins_from_piles(arr [][]int, k int) int {
	return _max_val_of_k_coins_from_piles(arr, k, 0)
}

func _max_val_of_k_coins_from_piles(arr [][]int, k int, start int) int {

	if start == len(arr) || k == 0 {
		return 0
	}

	//not take
	t := _max_val_of_k_coins_from_piles(arr, k, start+1)
	//take
	sum := 0
	for i := 0; i < min(k, len(arr[start])); i++ {
		sum += arr[start][i]
		t = max(t, sum+_max_val_of_k_coins_from_piles(arr, k-(i+1), start+1))
	}
	return t
}

//---------------------------------------------------------------------------

//-----------------------------memoization----------------------------------
func max_val_of_k_coins_from_piles_memo(arr [][]int, k int) int {
	n := len(arr)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			memo[i][j] = -1
		}
	}
	return _max_val_of_k_coins_from_piles_memo(arr, &memo, k, 0)
}

func _max_val_of_k_coins_from_piles_memo(arr [][]int, memo *[][]int, k int, start int) int {
	if start == len(arr) || k == 0 {
		return 0
	}

	if (*memo)[start][k] != -1 {
		return (*memo)[start][k]
	}

	//not take
	t := _max_val_of_k_coins_from_piles_memo(arr, memo, k, start+1)

	sum := 0
	for i := 0; i < min(k, len(arr[start])); i++ {
		sum += arr[start][i]
		t = max(t, sum+_max_val_of_k_coins_from_piles_memo(arr, memo, k-(i+1), start+1))
	}
	(*memo)[start][k] = t
	return (*memo)[start][k]
}

//-----------------------------memoization----------------------------------

func main() {
	fmt.Println(max_val_of_k_coins_from_piles_memo([][]int{
		{1, 100, 3}, {7, 8, 9},
	}, 2))
}
