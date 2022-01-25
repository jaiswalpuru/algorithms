package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//---------------------brute force TLE---------------------
func odd_subarray(arr []int) int {
	n := len(arr)

	cnt := 0
	for i := 0; i < n; i++ {
		sum := arr[i]
		if sum%2 != 0 {
			cnt++
		}
		for j := i + 1; j < n; j++ {
			sum += arr[j]
			if sum%2 != 0 {
				cnt++
			}
		}
	}
	return cnt
}

//---------------------------------------------------------------

const (
	mod = 1000000007
)

//the idea to see number of odd and even element
//odd+odd = even , even+odd = odd
func odd_sum_subarray_sum(arr []int) int {
	n := len(arr)
	odd, even := 0, 0
	sum := 0
	ans := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum%2 == 0 {
			ans += odd
			even++
		} else {
			ans += 1 + even
			odd++
		}
	}
	return ans % mod
}

func main() {
	fmt.Println(odd_sum_subarray_sum([]int{1, 3, 5}))
}
