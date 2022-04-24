package main

import "fmt"
import "sort"

func bsearch_lower_bound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	res := 0
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] >= target {
			res = len(arr) - mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func count_the_number_of_rectangles_containing_each_point(rec [][]int, points [][]int) []int {
	arr := make([][]int, 101)
	n := len(rec)
	for i := 0; i < n; i++ {
		arr[rec[i][1]] = append(arr[rec[i][1]], rec[i][0])
	}

	for i := 0; i < 101; i++ {
		sort.Ints(arr[i])
	}

	res := []int{}
	for i := 0; i < len(points); i++ {
		l, ht := points[i][0], points[i][1]
		sum := 0
		for j := ht; j < 101; j++ {
			sum += bsearch_lower_bound(arr[j], l)
		}
		res = append(res, sum)
	}
	return res
}

func main() {
	fmt.Println(count_the_number_of_rectangles_containing_each_point([][]int{
		{1, 2}, {2, 3}, {2, 5},
	}, [][]int{
		{2, 1}, {1, 4},
	}))
}
