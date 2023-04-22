package main

import "fmt"
import "sort"

//------------------using dfs---------------------
func stepping_numbers_dfs(l, h int) []int {
	res := []int{}
	mp := make(map[int]int)
	for i := 0; i <= 9; i++ {
		_backtrack(i, i, l, h, &mp)
	}
	for _, v := range mp {
		res = append(res, v)
	}
	sort.Ints(res)
	return res
}

func _backtrack(val, last, l, h int, mp *map[int]int) {
	if val > h {
		return
	}
	if val >= l {
		(*mp)[val] = val
	}

	if last == 0 {
		_backtrack(10*val+1, 1, l, h, mp)
	} else if last == 9 {
		_backtrack(10*val+8, 8, l, h, mp)
	} else {
		_backtrack(10*val+last-1, last-1, l, h, mp)
		_backtrack(10*val+last+1, last+1, l, h, mp)
	}
}

//------------------using dfs---------------------

//------------------using bfs---------------------
func stepping_numbers(l, h int) []int {
	res := []int{}
	if l == 0 {
		res = append(res, 0)
	}
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr <= h {
			if l <= curr {
				res = append(res, curr)
			}
			if curr%10 > 0 {
				q = append(q, 10*curr+curr%10-1)
			}
			if curr%10 < 9 {
				q = append(q, 10*curr+curr%10+1)
			}
		}
	}
	return res
}

//------------------using bfs---------------------

func main() {
	fmt.Println(stepping_numbers_dfs(0, 21))
}
