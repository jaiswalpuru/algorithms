package main

import "fmt"

var mod = int(1e9 + 7)

//-----------------------recursion--------------------------
func number_of_dice_rolls_with_target_sum(n, k, target int) int {
	cnt := 0
	recur(n, k, target, 0, &cnt)
	return cnt
}

func recur(n, k, target int, sum int, cnt *int) {
	if n == 0 {
		if target == sum {
			*cnt = (*cnt + 1) % mod
		}
		return
	}

	for i := 1; i <= k; i++ {
		if sum+i > target {
			continue
		}

		sum += i
		recur(n-1, k, target, sum, cnt)
		sum -= i
	}
}

//-----------------------recursion--------------------------

//-----------------------memo--------------------------
func number_of_dice_rolls_with_target_sum_memo(n, k, target int) int {
	cnt := 0
	memo := make(map[int]map[int]int)
	return _memo(n, k, target, 0, &cnt, &memo)
}

func _memo(n, k, target int, sum int, cnt *int, memo *map[int]map[int]int) int {
	if n == 0 {
		if target == sum {
			return 1
		}
		return 0
	}

	if _, ok := (*memo)[n]; ok {
		if _, k := (*memo)[n][sum]; k {
			return (*memo)[n][sum]
		}
	}

	s := 0
	for i := 1; i <= k; i++ {
		if sum+i > target {
			continue
		}
		sum += i
		s = (s + _memo(n-1, k, target, sum, cnt, memo)) % mod
		sum -= i
	}
	if _, ok := (*memo)[n]; !ok {
		(*memo)[n] = make(map[int]int)
	}
	(*memo)[n][sum] = s
	return (*memo)[n][sum]
}

//-----------------------memo--------------------------

func main() {
	fmt.Println(number_of_dice_rolls_with_target_sum_memo(1, 6, 3))
}
