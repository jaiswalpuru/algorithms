package main

import "fmt"

func split_array_into_consecutive_subsequences(arr []int) bool {
	freq := make(map[int]int)
	subs := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		freq[arr[i]]++
	}
	for i := 0; i < n; i++ {
		if freq[arr[i]] == 0 {
			continue
		}
		//check if a valid sequence exist with the last element which is arr[i]-1
		if subs[arr[i]-1] > 0 {
			subs[arr[i]-1]--
			subs[arr[i]]++
		} else if freq[arr[i]+1] > 0 && freq[arr[i]+2] > 0 {
			//if we want to start a new subsequence check if arr[i]+1 and arr[i]+2 both exist
			subs[arr[i]+2]++
			freq[arr[i]+1]--
			freq[arr[i]+2]--
		} else {
			return false
		}
		freq[arr[i]]--
	}
	return true
}

func main() {
	fmt.Println(split_array_into_consecutive_subsequences([]int{1, 2, 3, 3, 4, 5}))
}
