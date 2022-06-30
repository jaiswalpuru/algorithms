package main

import (
	"fmt"
	"math"
	"sort"
)

//--------------------------brute force--------------------------
func minimum_elements_to_equal_array_elements(arr []int) int {
	sort.Ints(arr)
	res := math.MaxInt64
	for i := 0; i < len(arr); i++ {
		t := 0
		for j := 0; j < len(arr); j++ {
			if i != j {
				t += int(math.Abs(float64(arr[i] - arr[j])))
			}
		}
		res = int(math.Min(float64(res), float64(t)))
	}
	return res
}

//-------------------------brute force---------------------------

//-------------------------efficient way---------------------------
func minimum_elements_to_equal_array_elements_eff(arr []int) int {
	sort.Ints(arr)
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += int(math.Abs(float64(arr[len(arr)/2]) - float64(arr[i])))
	}
	return sum
}

//-------------------------efficient way---------------------------

func main() {
	fmt.Println(minimum_elements_to_equal_array_elements_eff([]int{1, 2, 3}))
}
