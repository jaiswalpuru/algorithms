package main

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
	size := len(nums)
	res := int64(0)

	for i := 0; i < size; {
		if nums[i] == 0 {
			k := i
			zeros := 0
			for k < size {
				if nums[k] != 0 {
					break
				}
				zeros++
				k++
			}
			i = k
			res += int64(zeros) * int64(zeros+1) / 2
		} else {
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(number_of_zero_filled_subarrays([]int{1, 3, 0, 0, 2, 0, 0, 4}))
}
