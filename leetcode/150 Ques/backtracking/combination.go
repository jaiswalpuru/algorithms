/*
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
You may return the answer in any order.

Example 1:
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

Example 2:
Input: n = 1, k = 1
Output: [[1]]

*/

package main

import "fmt"

func _combination(n, k int, temp []int, res *[][]int, start int) {

	//fmt.Println("Called again")
	if len(temp) == k {
		//fmt.Println(start, temp, n)
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := start; i <= n; i++ {
		//fmt.Println("Before", i)
		_combination(n, k, append(temp, i), res, i+1)
		//fmt.Println("After", i)
	}
}

func combination(n, k int) [][]int {
	res := [][]int{}

	if n == 0 {
		return res
	}

	temp := []int{}
	_combination(n, k, temp, &res, 1)
	return res
}

func main() {
	fmt.Println(combination(4, 2))
	fmt.Println(combination(1, 1))
}
