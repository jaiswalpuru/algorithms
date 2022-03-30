package main

import "fmt"

//---------------------------brute force------------------------------
func songs_with_duration_divisible_by_60(arr []int) int {
	cnt := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (arr[i]+arr[j])%60 == 0 {
				cnt++
			}
		}
	}
	return cnt
}

//---------------------------brute force------------------------------

//----------------------------a better approach using hash_map-----------------
func songs_with_duration_divisible_by_60_eff(arr []int) int {
	hash_map := make(map[int]int)
	n := len(arr)
	cnt := 0
	for i := 0; i < n; i++ {
		if arr[i]%60 == 0 {
			cnt += hash_map[0]
		} else {
			cnt += hash_map[60-arr[i]%60]
		}
		hash_map[arr[i]%60]++
	}
	return cnt
}

//----------------------------a better approach using hash_map-----------------

func main() {
	fmt.Println(songs_with_duration_divisible_by_60([]int{30, 20, 150, 100, 40}))
}

//purujaiswal12#
