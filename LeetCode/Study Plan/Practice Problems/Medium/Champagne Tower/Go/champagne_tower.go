package main

import (
	"fmt"
	"math"
)

//-------------------------------recursion (TLE)-------------------------------
func champagne_tower_recursion(poured, row, glass int) float64 {
	return _champagne_tower_recursion(poured, row, glass)
}

func _champagne_tower_recursion(poured, row, glass int) float64 {
	if row < 0 || glass < 0 || glass > row {
		return 0
	}

	if row == 0 && glass == 0 {
		return float64(poured)
	}

	left := (_champagne_tower_recursion(poured, row-1, glass-1) - 1) / 2.00
	if left < 0 {
		left = 0
	}

	right := (_champagne_tower_recursion(poured, row-1, glass) - 1) / 2.00
	if right < 0 {
		right = 0
	}

	return left + right
}

//------------------------------------------------------------------------------------

//-------------------------------------------dp-----------------------------------------
func champagne_tower_dp(poured, row, glass int) float64 {
	dp := make([][]float64, 102)
	for i := 0; i < 102; i++ {
		dp[i] = make([]float64, 102)
	}
	dp[0][0] = float64(poured)
	for i := 0; i <= row; i++ {
		for j := 0; j <= i; j++ {
			poured_left := float64(dp[i][j]-1.0) / 2.00
			if poured_left > 0 {
				dp[i+1][j] += poured_left
				dp[i+1][j+1] += poured_left
			}
		}
	}
	return math.Min(1, dp[row][glass])
}

//------------------------------------------------------------------------------------

func main() {
	fmt.Println(champagne_tower_dp(2, 1, 1))
}
