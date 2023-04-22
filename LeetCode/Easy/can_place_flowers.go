/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in
adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.
*/

package main

import "fmt"

func flower_bed(arr []int, n int) bool {
	size := len(arr)

	if arr[0] == 0 {
		if size > 1 {
			if arr[1] == 0 {
				arr[0] = 1
				n--
			}
		}
	}

	for i := 1; i < size-1; i++ {
		if arr[i] == 0 {
			if arr[i-1] == arr[i] && arr[i+1] == arr[i] {
				arr[i] = 1
				n--
			}
		}
	}

	if arr[size-1] == 0 && n > 0 && size > 1 {
		if arr[size-2] == 0 {
			arr[size-1] = 1
			n--
		}
	}

	if n <= 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(flower_bed([]int{0, 0, 0, 0}, 1))
	// fmt.Println(flower_bed([]int{1, 0, 0, 0, 1}, 2))
	// fmt.Println(flower_bed([]int{0, 0, 1, 0, 1}, 1))
	// fmt.Println(flower_bed([]int{1, 0, 0, 0, 1, 0, 0}, 2))
}
