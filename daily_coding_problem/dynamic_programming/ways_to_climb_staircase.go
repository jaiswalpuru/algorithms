/*
There exists a staircase with n steps which you can climb up either 1 or 2 steps at a time. Given n, write the
function which returns the number of unique ways you can climb the staircase.

eg : n=4
1,1,1,1
2,1,1
1,2,1
1,1,2
2,2
*/

package main

import "fmt"

//recursive
func num_ways_recursive(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		return num_ways_recursive(n-1) + num_ways_recursive(n-2)
	}
}

//bottom up
func num_ways(n int) int {

	if n <= 0 {
		return 0
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func main() {
	fmt.Println(num_ways(4))
	fmt.Println(num_ways_recursive(4))

	lets say we are given an array which consists of steps we can take
	fmt.Println(num_ways_mod(4, []int{1, 3, 5}))
	fmt.Println(num_ways_mod_dp(7, []int{2, 3, 6, 7}))
}

func num_ways_mod_dp(n int, x []int) int {
	arr := make([]int, n+1)
	m := len(x)
	arr[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			if i-x[j] >= 0 {
				arr[i] += arr[i-x[j]]
			}
		}
	}
	return arr[n]
}

func num_ways_mod(n int, arr []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		sum := 0
		for _, v := range arr {
			sum += num_ways_mod(n-v, arr)
		}
		return sum
	}
}
