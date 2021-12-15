/*
	Given an array of integers where every integer occurs three times except for one integer,
	which only occurs once, find and return the non-duplicated integer.

	For example, given [6, 1, 3, 3, 3, 6, 6], return 1. Given [13, 19, 13, 13], return 19.

	Do this in O(N) time and O(1) space.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{6, 1, 3, 3, 3, 6, 6}
	//arr = []int{13, 13, 16, 13}
	three_n, three_n1, three_n2 := int(math.MaxInt8), 0, 0

	for i := 0; i < len(arr); i++ {

		common_n := arr[i] & three_n
		common_n1 := arr[i] & three_n1
		common_n2 := arr[i] & three_n2

		three_n &= (^common_n)
		three_n1 |= (common_n)

		three_n1 &= (^common_n1)
		three_n2 |= (common_n1)

		three_n2 &= (^common_n2)
		three_n |= (common_n2)
	}

	fmt.Println(three_n1)
}
