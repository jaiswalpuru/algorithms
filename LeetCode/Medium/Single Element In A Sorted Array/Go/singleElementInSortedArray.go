package main

import "fmt"

//------------O(N)---------------------
func singleNonDuplicateBrute(nums []int) int {
	xor := nums[0]
	size := len(nums)

	for i := 1; i < size; i++ {
		xor = xor ^ nums[i]
	}

	return xor
}

//------------O(N)---------------------

//--------------O(logn)-----------------
func singleNonDuplicate(nums []int) int {
	size := len(nums)
	low := 0
	high := size - 1

	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid+1] {
			if (high-mid)%2 == 0 {
				low = mid + 2
			} else {
				high = mid - 1
			}
		} else if nums[mid] == nums[mid-1] {
			if (high-mid)%2 == 0 {
				high = mid - 2
			} else {
				low = mid + 1
			}
		} else {
			return nums[mid]
		}
	}

	return nums[low]
}

//--------------O(logn)-----------------

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
