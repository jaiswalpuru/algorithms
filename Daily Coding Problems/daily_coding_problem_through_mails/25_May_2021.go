/*
	Given a list of integers S and a target number k, write a function that returns a subset of S that adds up to k.
	If such a subset cannot be made, then return null.

	Integers can appear more than once in the list. You may assume all numbers in the list are positive.

	For example, given S = [12, 1, 61, 5, 9, 2] and k = 24, return [12, 9, 2, 1] since it sums up to 24.
*/

package main

import (
	"fmt"
	"sort"
)

var dp = make([][]int, 62)

func display(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func print_all_susbset(arr []int, i, sum int, p []int) {

	//reached end and sum is non zero only print when arr[0] == sum or dp[0][sum] == 1
	if i == 0 && sum != 0 && dp[arr[0]][sum] == 1 {
		p = append(p, arr[i])
		if arr[i] == sum {
			display(p)
		}
		return
	}

	//if sum becomes zero
	if i == 0 && sum == 0 {
		display(p)
		return
	}

	//if the given sum can be acheieved by ignoring the current element
	if dp[arr[i-1]][sum] == 1 {
		b := p
		print_all_susbset(arr, i-1, sum, b)
	}

	//if the given sum can be considered after taking the current element into consideration
	if sum >= arr[i] && dp[arr[i-1]][sum-arr[i]] == 1 {
		fmt.Println(arr, i, sum, p, arr[i], sum-arr[i])
		p = append(p, arr[i])
		print_all_susbset(arr, i-1, sum-arr[i], p)
	}
}

func main() {
	arr := []int{12, 1, 61, 5, 9, 2}
	k := 24

	sort.Ints(arr)
	n := len(arr)

	for i := 0; i < n; i++ {
		dp[arr[i]] = make([]int, k+1)
	}

	for i := 0; i < n; i++ {
		dp[arr[i]][0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 1; j < (k + 1); j++ {
			if i != 0 {
				if arr[i] > j {
					dp[arr[i]][j] = dp[arr[i-1]][j]
				} else {
					if dp[arr[i-1]][j] == 1 || dp[arr[i-1]][j-arr[i]] == 1 {
						dp[arr[i]][j] = 1
					}
				}
			} else {
				if j == arr[i] {
					dp[arr[i]][j] = 1
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if dp[arr[i]] != nil {
			fmt.Println(arr[i], " ", dp[arr[i]])
		}
	}

	if dp[arr[n-1]][k] == 1 {
		fmt.Printf("Subset sum exists for k=%d and arr = []int{%v}\n", k, arr)
		//To find the subset we need to trace back the path
		ans := []int{}
		print_all_susbset(arr, n-1, k, ans)
	} else {
		fmt.Println("Susbset doesn't exists")
	}
}
