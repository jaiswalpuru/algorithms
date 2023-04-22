package main

import "fmt"

//------------------------------brute force-------------------------------
func sum_of_even_number_after_queries(nums []int, queries [][]int) []int {
	n := len(nums)
	m := len(queries)
	res := []int{}
	for i := 0; i < m; i++ {
		val := queries[i]
		nums[val[1]] += val[0]
		sum := 0
		for j := 0; j < n; j++ {
			if nums[j]%2 == 0 {
				sum += nums[j]
			}
		}
		res = append(res, sum)
	}
	return res
}

//------------------------------brute force-------------------------------

//------------------------------optimized code----------------------------
func sum_of_even_number_after_queries_optimized(nums []int, queries [][]int) []int {
	n := len(nums)
	m := len(queries)
	sum := 0
	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			sum += nums[i]
		}
	}
	res := []int{}
	for i := 0; i < m; i++ {
		val, j := queries[i][0], queries[i][1]
		if nums[j]%2 == 0 {
			sum -= nums[j]
		}
		nums[j] += val
		if nums[j]%2 == 0 {
			sum += nums[j]
		}
		res = append(res, sum)
	}
	return res
}

//------------------------------optimized code----------------------------

func main() {
	fmt.Println(sum_of_even_number_after_queries_optimized([]int{1, 2, 3, 4}, [][]int{
		{1, 0}, {-3, 1}, {-4, 0}, {2, 3},
	}))
}
