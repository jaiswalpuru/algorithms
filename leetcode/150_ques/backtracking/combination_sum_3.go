/*
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:
Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations.
The list must not contain the same combination twice, and the combinations may be returned in any order.

Example 1:
Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.

Example 2:
Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.

Example 3:
Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.

Example 4:
Input: k = 3, n = 2
Output: []
Explanation: There are no valid combinations.

Example 5:
Input: k = 9, n = 45
Output: [[1,2,3,4,5,6,7,8,9]]
Explanation:
1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 = 45
There are no other valid combinations.
*/

package main

import "fmt"

func _combination_3(arr []int, res *[][]int, k, n int, temp []int, ind int) {

	if len(temp) == k && n == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := ind; i < len(arr); i++ {
		if n-arr[i] < 0 {
			break
		}
		temp = append(temp, arr[i])
		_combination_3(arr, res, k, n-arr[i], temp, i+1)
		temp = temp[:len(temp)-1]
	}
}

func combination_3(k, n int) [][]int {
	res := [][]int{}
	temp := []int{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	_combination_3(arr, &res, k, n, temp, 0)
	return res
}

func main() {
	fmt.Println(combination_3(3, 7))
	fmt.Println(combination_3(3, 9))
	fmt.Println(combination_3(4, 1))
	fmt.Println(combination_3(3, 2))
	fmt.Println(combination_3(9, 45))
}
