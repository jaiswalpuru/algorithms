package main

import (
	"fmt"
)

//----------------------brute force------------------------
func count_smaller_number_after_self(arr []int) []int {
	n := len(arr)
	smaller := make([]int, n)
	for i := 0; i < n; i++ {
		val := arr[i]
		cnt := 0
		for j := i + 1; j < n; j++ {
			if arr[j] < val {
				cnt++
			}
		}
		smaller[i] = cnt
	}
	return smaller
}

//----------------------brute force------------------------
//----------------------efficient approach------------------------

func count_smaller_number_after_self_eff(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	merge_sort(arr, indices, &result, 0, n)
	return result
}

func merge_sort(arr []int, indices []int, res *[]int, left, right int) {
	if left >= right-1 {
		return
	}
	mid := (left + right) / 2
	merge_sort(arr, indices, res, left, mid)
	merge_sort(arr, indices, res, mid, right)
	merge(arr, indices, left, right, mid, res)
}

func merge(arr []int, indices []int, left, right, mid int, res *[]int) {
	i, j := left, mid
	temp := []int{}
	for i < mid && j < right {
		if arr[indices[i]] <= arr[indices[j]] {
			(*res)[indices[i]] += j - mid
			temp = append(temp, indices[i])
			i++
		} else {
			temp = append(temp, indices[j])
			j++
		}
	}

	for i < mid {
		(*res)[indices[i]] += j - mid
		temp = append(temp, indices[i])
		i++
	}

	for j < right {
		temp = append(temp, indices[j])
		j++
	}

	for k := left; k < right; k++ {
		indices[k] = temp[k-left]
	}
}

//----------------------efficient approach------------------------

func main() {
	fmt.Println(count_smaller_number_after_self_eff([]int{5, 2, 6, 1}))
}
