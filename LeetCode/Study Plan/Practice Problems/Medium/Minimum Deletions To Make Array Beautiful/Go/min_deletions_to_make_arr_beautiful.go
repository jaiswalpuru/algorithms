package main

import "fmt"

func min_deletions(nums []int) int {
	n := len(nums)
	del := 0
	for i := 0; i < n-1; i++ {
		if (i-del)%2 == 0 {
			if nums[i] == nums[i+1] {
				nums[i] = -1
				del++
			}
		}
	}

	if (n-del)%2 == 0 {
		return del
	} else {
		for i := 0; i < n; i++ {
			if nums[i] != -1 {
				nums[i] = -1
				del++
				if (n-del)%2 == 0 {
					break
				}
			}
		}
	}

	return del
}

func main() {
	fmt.Println(min_deletions([]int{1, 1, 2, 3, 5}))
}
