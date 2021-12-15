/*

	We can determine how "out of order" an array A is by counting the number of inversions it has.
	Two elements A[i] and A[j] form an inversion if A[i] > A[j] but i < j.
	That is, a smaller element appears after a larger element.

	Given an array, count the number of inversions it has. Do this faster than O(N^2) time.

	You may assume each element in the array is distinct.

	For example, a sorted list has zero inversions. The array [2, 4, 1, 3, 5] has
	three inversions: (2, 1), (4, 1), and (4, 3).
	The array [5, 4, 3, 2, 1] has ten inversions: every distinct pair forms an inversion.

*/

package main

import "fmt"

func count_inversion_helper(arr []int) (int, []int) {
	if len(arr) <= 1 {
		return 0, arr
	}
	mid := len(arr) / 2
	left_cnt, left_sorted_arr := count_inversion_helper(arr[:mid])
	right_cnt, right_sorted_arr := count_inversion_helper(arr[mid:])
	between_cnt, sorted_arr := merge_count(left_sorted_arr, right_sorted_arr)
	return left_cnt + right_cnt + between_cnt, sorted_arr
}

func merge_count(a, b []int) (int, []int) {
	cnt := 0
	sorted_arr := []int{}

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted_arr = append(sorted_arr, a[i])
			i++
		} else if a[i] > b[j] {
			sorted_arr = append(sorted_arr, b[j])
			cnt += len(a) - i
			j++
		}
	}

	if i < len(a) {
		sorted_arr = append(sorted_arr, a[i:]...)
	}
	if j < len(b) {
		sorted_arr = append(sorted_arr, b[j:]...)
	}

	return cnt, sorted_arr
}

func count_inversions(arr []int) int {
	cnt, _ := count_inversion_helper(arr)
	return cnt
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(count_inversions(arr))
}
