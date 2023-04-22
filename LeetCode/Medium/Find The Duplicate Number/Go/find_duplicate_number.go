package main

import (
	"fmt"
	"sort"
)

//--------------------using sorting---------------------------- (dont do this as it is not allowed but will run fine)
func find_duplicate_number_sort(arr []int) int {
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		if arr[i] == arr[i+1] {
			return arr[i]
		}
	}
	return -1
}

//--------------------using sorting---------------------------- (refer LC)
func find_duplicate_number(arr []int) int {
	low, high := 1, len(arr)-1 //low and high represents the range of values
	duplicate := -1
	for low <= high {
		mid := (low + high) / 2

		cnt := 0

		//counting the number of elements which are smaller than mid
		for _, v := range arr {
			if v <= mid {
				cnt++
			}
		}

		if cnt > mid {
			duplicate = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return duplicate
}

func main() {
	fmt.Println(find_duplicate_number([]int{3, 1, 3, 4, 2}))
}
