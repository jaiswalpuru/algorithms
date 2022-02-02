package main

import (
	"fmt"
	"math"
)

//----------------------starting both the persons from 0,0--------------------------------------------------------
func cherry_pick(grid [][]int) int {
	return _cherry_pick(grid, 0, 0, 0)
}

func _cherry_pick(grid [][]int, r1, c1, c2 int) int {
	r2 := r1 + c1 - c2 //**important**
	if r1 == len(grid) || r2 == len(grid) || c1 == len(grid) || c2 == len(grid) || (grid)[r1][c1] == -1 ||
		(grid)[r2][c2] == -1 {
		return -1e8
	} else if r1 == len(grid)-1 && c1 == len(grid)-1 {
		return (grid)[r1][c1]
	} else {
		ans := (grid)[r1][c1]
		if c1 != c2 {
			ans += (grid)[r2][c2]
		}
		ans += max(max(_cherry_pick(grid, r1, c1+1, c2+1), _cherry_pick(grid, r1+1, c1, c2+1)),
			max(_cherry_pick(grid, r1, c1+1, c2), _cherry_pick(grid, r1+1, c1, c2)))
		return ans
	}
}

//----------------------------------------------------------------------------------------------------------------

//--------------------------------------------memoization---------------------------------------------------
func cherry_pick_memo(grid [][]int) int {
	n := len(grid)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = math.MinInt64
			}
		}
	}

	return max(0, _cherry_pick_memo(grid, &dp, 0, 0, 0))
}

func _cherry_pick_memo(grid [][]int, dp *[][][]int, r1, c1, c2 int) int {
	r2 := r1 + c1 - c2
	if r1 == len(grid) || r2 == len(grid) || c1 == len(grid) || c2 == len(grid) || grid[r2][c1] == -1 ||
		grid[r1][c2] == -1 {
		return -1e8
	} else if (*dp)[r1][c1][c2] != math.MinInt64 {
		return (*dp)[r1][c1][c2]
	} else if r1 == len(grid)-1 && c1 == len(grid)-1 {
		return grid[r1][c1]
	} else {
		ans := grid[r1][c1]
		if c1 != c2 {
			ans += grid[r2][c2]
		}
		ans += max(max(_cherry_pick_memo(grid, dp, r1+1, c1, c2), _cherry_pick_memo(grid, dp, r1, c1+1, c2)),
			max(_cherry_pick_memo(grid, dp, r1+1, c1, c2+1), _cherry_pick_memo(grid, dp, r1, c1+1, c2+1)))
		(*dp)[r1][c1][c2] = ans
		return (*dp)[r1][c1][c2]
	}
}

//--------------------------------------------memoization-----------------------------------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(cherry_pick_memo([][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1},
	}))
}
