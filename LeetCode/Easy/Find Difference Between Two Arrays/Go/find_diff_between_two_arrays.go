package main

import "fmt"

func find_diff_between_two_arr(nums1, nums2 []int) [][]int {
	res := [][]int{}
	h1 := make(map[int]bool)
	h2 := make(map[int]bool)
	n, m := len(nums1), len(nums2)
	for i := 0; i < n; i++ {
		h1[nums1[i]] = true
	}
	visited := make(map[int]bool)
	for i := 0; i < m; i++ {
		h2[nums2[i]] = true
	}

	t := []int{}
	for i := 0; i < n; i++ {
		if !h2[nums1[i]] && !visited[nums1[i]] {
			t = append(t, nums1[i])
			visited[nums1[i]] = true

		}
	}

	res = append(res, t)
	t = []int{}
	visited = make(map[int]bool)
	for i := 0; i < m; i++ {
		if !h1[nums2[i]] && !visited[nums2[i]] {
			t = append(t, nums2[i])
			visited[nums2[i]] = true
		}
	}
	res = append(res, t)
	return res
}

func main() {
	fmt.Println(find_diff_between_two_arr([]int{1, 2, 3}, []int{2, 4, 6}))
}
