/*
Given a real number n, find the square root of n. For example, given n = 9, return 3.
*/

package main

import "fmt"

//TC : O(sqrt(n))
func square_root(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	res, i := 1, 1
	for res <= n {
		i++
		res = i * i
	}
	return i - 1
}

/*
The idea is to use binary search.
The value i*i is increasing monotonically so we can use binary search.
*/
func square_root_method_two(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	lb, ub := 0, n
	ans := 0
	for lb <= ub {
		mid := (lb + ub) / 2
		if mid*mid == n {
			return mid
		}

		if mid <= n/mid {
			lb = mid + 1
			ans = mid
		} else {
			ub = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(square_root_method_two(9))
}
