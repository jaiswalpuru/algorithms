package main

import "fmt"

var MAX_COST = int(1e6 + 1)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//-------------------------------brute force --------------------------------------
func paint_houses_three(houses []int, cost [][]int, m, n int, target int) int {
	res := _recur(houses, cost, target, 0, 0, 0)
	if res == MAX_COST {
		return -1
	}
	return res
}

func _recur(houses []int, cost [][]int, target int, ind, neig_cnt, prev_house_color int) int {
	if ind == len(houses) {
		if neig_cnt == target {
			return 0
		}
		return MAX_COST
	}

	if neig_cnt > target {
		return MAX_COST
	}

	min_cost := MAX_COST

	if houses[ind] != 0 {
		new_neig_cnt := neig_cnt
		if houses[ind] != prev_house_color {
			new_neig_cnt = neig_cnt + 1
		}
		min_cost = _recur(houses, cost, target, ind+1, new_neig_cnt, houses[ind])
	} else {
		colors := len(cost[0])
		for i := 1; i <= colors; i++ {
			new_neig_cnt := neig_cnt
			if prev_house_color != i {
				new_neig_cnt += 1
			}
			curr_cost := cost[ind][i-1] + _recur(houses, cost, target, ind+1, new_neig_cnt, i)
			min_cost = min(min_cost, curr_cost)
		}
	}
	return min_cost
}

//-------------------------------brute force --------------------------------------

//-------------------------------memoization--------------------------------------
func paint_houses_eff(houses []int, cost [][]int, m, n, target int) int {
	memo := make([][][21]int, 100)
	for i := 0; i < 100; i++ {
		memo[i] = make([][21]int, 100)
		for j := 0; j < 100; j++ {
			for k := 0; k < 21; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	res := _memo(houses, cost, target, 0, 0, 0, &memo)
	if res == MAX_COST {
		return -1
	}
	return res
}

func _memo(houses []int, cost [][]int, target, ind, neig_cnt, prev_color int, memo *[][][21]int) int {
	if ind == len(houses) {
		if neig_cnt == target {
			return 0
		}
		return MAX_COST
	}

	if neig_cnt > target {
		return MAX_COST
	}

	if (*memo)[ind][neig_cnt][prev_color] != -1 {
		return (*memo)[ind][neig_cnt][prev_color]
	}

	min_cost := MAX_COST
	if houses[ind] != 0 {
		new_neig_cnt := neig_cnt
		if houses[ind] != prev_color {
			new_neig_cnt += 1
		}
		min_cost = _memo(houses, cost, target, ind+1, new_neig_cnt, houses[ind], memo)
	} else {
		colors := len(cost[0])
		for i := 1; i <= colors; i++ {
			new_neig_cnt := neig_cnt
			if i != prev_color {
				new_neig_cnt += 1
			}
			curr_cost := cost[ind][i-1] + _memo(houses, cost, target, ind+1, new_neig_cnt, i, memo)
			min_cost = min(min_cost, curr_cost)
		}
	}
	(*memo)[ind][neig_cnt][prev_color] = min_cost
	return min_cost
}

//-------------------------------memoization--------------------------------------

func main() {
	fmt.Println(paint_houses_eff([]int{0, 0, 0, 0, 0}, [][]int{
		{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1},
	}, 5, 2, 3))
}
