package main

import "fmt"
import "math"

func max(a,b int)int {
    if a>b{
        return a
    }
    return b
}

func min(a, b int) int {
    if a>b {
        return b
    }
    return a
}

func fair_distribution_of_cookies(arr []int, k int) int {
	cook := make([]int, k)
	res := math.MaxInt64
	backtrack(arr, &cook, 0, k, &res)
	return res
}

func backtrack(arr []int, cook *[]int, ind, k int, ans *int) {
	if ind == len(arr) {
		temp := math.MinInt64
		for i := 0; i < k; i++ {
			temp = max(temp, (*cook)[i])
		}
		*ans = min(*ans, temp)
		return
	}

	for i := 0; i < k; i++ {
		(*cook)[i] += arr[ind]
		backtrack(arr, cook, ind+1, k, ans)
		(*cook)[i] -= arr[ind]
	}
}

func main() {
	fmt.Println(fair_distribution_of_cookies([]int{8, 15, 10, 20, 8}, 2))
}
