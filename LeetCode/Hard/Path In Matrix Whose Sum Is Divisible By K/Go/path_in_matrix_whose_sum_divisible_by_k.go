package main

import "fmt"

var mod = int(1e9 + 7)

//---------------------------brute force-------------------------------
func path_in_matrix_whose_sum_divisible_by_k(grid [][]int, k int) int {
	cnt := 0
	n, m := len(grid), len(grid[0])
	recur(grid, k, &cnt, n-1, m-1, 0)
	return cnt
}

func recur(grid [][]int, k int, cnt *int, n, m int, sum int) {
	if n == 0 && m == 0 {
		if (sum+grid[n][m])%k == 0 {
			*cnt++
		}
		return
	}
	if n < 0 || m < 0 {
		return
	}
	temp := sum + grid[n][m]
	recur(grid, k, cnt, n-1, m, temp)
	recur(grid, k, cnt, n, m-1, temp)
}

//---------------------------brute force-------------------------------

//---------------------------memo-------------------------------
func path_in_matrix_whose_sum_divisible_by_k_memo(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	mem := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			mem[i][j] = make([]int, k+1)
			for l := 0; l <= k; l++ {
				mem[i][j][l] = -1
			}
		}
	}
	return memo(grid, k, n-1, m-1, 0, &mem) % mod
}

func memo(grid [][]int, k int, n, m int, sum int, mem *[][][]int) int {
	if n == 0 && m == 0 {
		if (sum+grid[n][m])%k == 0 {
			return 1
		}
		return 0
	}
	if n < 0 || m < 0 {
		return 0
	}

	if (*mem)[n][m][sum] != -1 {
		return (*mem)[n][m][sum] % mod
	}

	up := memo(grid, k, n-1, m, (sum+grid[n][m])%k, mem)
	left := memo(grid, k, n, m-1, (sum+grid[n][m])%k, mem)
	(*mem)[n][m][sum] = (up + left) % mod
	return (*mem)[n][m][sum]
}

//---------------------------memo-------------------------------

func main() {
	fmt.Println(path_in_matrix_whose_sum_divisible_by_k_memo([][]int{
		{5, 2, 4}, {3, 0, 5}, {0, 7, 2},
	}, 3))
}
