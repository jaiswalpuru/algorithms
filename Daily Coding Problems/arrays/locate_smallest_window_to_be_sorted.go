package main

import (
	"fmt"
	"math"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func windowMethodTwo(arr []int) (int, int) {
	left, right := 0, 0
	maxSeen, minSeen := int(math.MinInt32), int(math.MaxInt32)

	for i := 0; i < len(arr); i++ {
		maxSeen = max(maxSeen, arr[i])
		if maxSeen > arr[i] {
			right = i
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		minSeen = min(minSeen, arr[i])
		if arr[i] > minSeen {
			left = i
		}
	}
	return left, right
}

// O(nlogn)
func windowMethodOne(arr, temp []int) (int, int) {

	left, right := 0, 0
	sort.Ints(temp)
	for i := 0; i < len(arr); i++ {
		if arr[i] != temp[i] && left == 0 {
			left = i
		} else if arr[i] != temp[i] {
			right = i
		}
	}
	return left, right

}

func main() {
	arr := []int{3, 7, 5, 6, 9}
	temp := []int{3, 7, 5, 6, 9}
	l, r := windowMethodOne(arr, temp)
	fmt.Println(l, r)
	l, r = windowMethodTwo(arr)
	fmt.Println(l, r)
}
