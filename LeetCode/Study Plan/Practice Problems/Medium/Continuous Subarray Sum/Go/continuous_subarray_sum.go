package main

import "fmt"

//----------------------brute force---------------------
func continuous_subarray_sum(arr []int, k int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		sum := arr[i]
		for j := i + 1; j < n; j++ {
			sum += arr[j]
			if sum%k == 0 {
				return true
			}
		}
	}
	return false
}

//----------------------brute force---------------------

//----------------------efficient approach---------------------
func continuous_subarray_sum_eff(arr []int, k int) bool {
	mp := make(map[int]int) //map of remainder and index
	mp[0] = -1              //so that if we see 0 as a remainder in future in future we will not assume that there are atleast two elements
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		remain := sum % k
		if v, ok := mp[remain]; !ok {
			mp[remain] = i
		} else {
			if i-v > 1 {
				return true
			}
		}
	}
	return false
}

//----------------------efficient approach---------------------

func main() {
	fmt.Println(continuous_subarray_sum_eff([]int{23, 2, 4, 6, 7}, 6))
}
