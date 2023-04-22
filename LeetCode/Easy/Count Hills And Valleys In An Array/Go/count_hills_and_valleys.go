package main

import "fmt"

func count_hill_valley(arr []int) int {
	n := len(arr)
	cnt := 0
	for i := 1; i < n-1; i++ {
		if i == 1 || arr[i] != arr[i-1] {
			var j, k int
			for j = i - 1; j >= 0; j-- {
				if arr[j] != arr[i] {
					break
				}
			}
			for k = i + 1; k < n; k++ {
				if arr[k] != arr[i] {
					break
				}
			}

			if j >= 0 && k < n && arr[i] > arr[j] && arr[i] > arr[k] {
				cnt++
			}
			if j >= 0 && k < n && arr[i] < arr[j] && arr[i] < arr[k] {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(count_hill_valley([]int{2, 4, 1, 1, 6, 5}))
}
