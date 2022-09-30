package main

import (
	"fmt"
	"sort"
)

//---------------------brute force---------------------
func maximum_population_year(logs [][]int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})
	n := len(logs)
	cnt := 1
	res := logs[0][0]
	for i := 0; i < n; i++ {
		t := 0
		for j := 0; j < n; j++ {
			if logs[i][0] >= logs[j][0] && logs[i][0] < logs[j][1] {
				t++
			}
		}
		if t > cnt {
			cnt = t
			res = logs[i][0]
		}
	}
	return res
}

//---------------------brute force---------------------

func main() {
	fmt.Println(maximum_population_year([][]int{
		{1993, 1999}, {2000, 2010},
	}))
}
