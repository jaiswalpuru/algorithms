package main

import "fmt"

//---------------------------brute force---------------------
func subarray_sum_divisible_by_k(arr []int, k int) int {
	n := len(arr)
	cnt := 0
	for i := 0; i < n; i++ {
		sum := arr[i]
		if sum%k == 0 {
			cnt++
		}
		for j := i + 1; j < n; j++ {
			sum += arr[j]
			if sum%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}

//---------------------------brute force---------------------

//------------------------efficient approach------------------------
func subarray_sum_divisible_by_k_eff(arr []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	n := len(arr)
	cnt := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum = ((sum+arr[i])%k + k) % k // for negative elements
		cnt += mp[sum]
		mp[sum]++
	}
	return cnt
}

//------------------------efficient approach-----------------------

func main() {
	fmt.Println(subarray_sum_divisible_by_k_eff([]int{4, 5, 0, -2, -3, 1}, 5))
}
