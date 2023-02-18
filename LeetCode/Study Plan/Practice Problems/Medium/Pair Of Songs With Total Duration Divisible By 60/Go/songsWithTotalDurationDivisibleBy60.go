package main

import "fmt"

//---------------------------brute force------------------------------
func songsWithDurationDivisibleBy60Brute(arr []int) int {
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
func numPairsDivisibleBy60(time []int) int {
	remainders := make([]int, 60)
	size := len(time)
	pairs := 0

	for i := 0; i < size; i++ {
		if time[i]%60 == 0 {
			pairs += remainders[0]
		} else {
			pairs += remainders[60-time[i]%60]
		}
		remainders[time[i]%60]++
	}

	return pairs
}

//----------------------------a better approach using hash_map-----------------

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
}

//purujaiswal12#
