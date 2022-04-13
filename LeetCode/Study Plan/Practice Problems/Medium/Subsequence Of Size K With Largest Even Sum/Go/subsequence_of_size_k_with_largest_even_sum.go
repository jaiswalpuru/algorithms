package main

import "fmt"
import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------------------brute force dont use---------------------------------------
//just generating all the length three subsequence and comparing the sum
func subsequence_of_size_k_with_largest_even_sum(arr []int, k int) int {
	sum := 0
	_subsequence_of_size_k_with_largest_even_sum(arr, k, 0, &sum, 0, []int{})
	return sum
}

func _subsequence_of_size_k_with_largest_even_sum(arr []int, k int, ind int, sum *int, temp int, set []int) {
	if ind == len(arr) {
		if len(set) == k {
			if temp%2 == 0 {
				*sum = max(*sum, temp)
			}
		}
		return
	}

	t := append(set, arr[ind])
	if len(t) > k {
		return
	}
	_subsequence_of_size_k_with_largest_even_sum(arr, k, ind+1, sum, temp+arr[ind], t)
	t = t[:len(t)-1]
	_subsequence_of_size_k_with_largest_even_sum(arr, k, ind+1, sum, temp, t)
}

//-------------------------------brute force dont use------------------------------------------

//-----------------------------greedy approach------------------------------------
func subsequence_of_size_k_with_largest_even_sum_greedy(arr []int, k int) int {
	odd, even := []int{}, []int{}
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			even = append(even, arr[i])
		} else {
			odd = append(odd, arr[i])
		}
	}

	sort.Ints(odd)
	sort.Ints(even)
	sum := -1
	even_prefix, odd_prefix := []int{0}, []int{0}
	for i := 0; i < len(even); i++ {
		even_prefix = append(even_prefix, even_prefix[len(even_prefix)-1]+even[len(even)-i-1])
	}

	for i := 0; i < len(odd); i++ {
		odd_prefix = append(odd_prefix, odd_prefix[len(odd_prefix)-1]+odd[len(odd)-1-i])
	}
	for i := 0; i <= len(odd); i += 2 {
		if i >= 0 && k >= i && k-i <= len(even) {
			sum = max(sum, odd_prefix[i]+even_prefix[k-i])
		}
	}

	return sum
}

//-----------------------------greedy approach------------------------------------

func main() {
	fmt.Println(subsequence_of_size_k_with_largest_even_sum_greedy([]int{1, 1, 3, 5, 4}, 3))
}
