/*
Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous) subarray of arr.
Since the answer may be large, return the answer modulo 109 + 7.
*/

package main

const (
	mod = 1000000007
)

type Pair struct {
	val, cnt int
}

func subarray_sum(arr []int) int {
	n := len(arr)

	left := make([]int, n)
	right := make([]int, n)

	l_stack := []Pair{}
	r_stack := []Pair{}
	//left
	for i := 0; i < n; i++ {
		curr := arr[i]
		cnt := 1
		for len(l_stack) > 0 && l_stack[len(l_stack)-1].val > curr {
			cnt += l_stack[len(l_stack)-1].cnt
			l_stack = l_stack[:len(l_stack)-1]
		}
		l_stack = append(l_stack, Pair{curr, cnt})
		left[i] = cnt
	}

	//right
	for i := n - 1; i >= 0; i-- {
		curr := arr[i]
		cnt := 1
		for len(r_stack) > 0 && r_stack[len(r_stack)-1].val >= curr {
			cnt += r_stack[len(r_stack)-1].cnt
			r_stack = r_stack[:len(r_stack)-1]
		}
		r_stack = append(r_stack, Pair{curr, cnt})
		right[i] = cnt
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = (ans + left[i]*right[i]*arr[i]) % mod
	}

	return ans
}

func main() {

}
