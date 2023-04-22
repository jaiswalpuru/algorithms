package main

import "fmt"

func sort_array_by_moving_to_empty_space(nums []int) int {
	n := len(nums)
	parent_one := make([]int, n)
	size_one := make([]int, n)
	parent_two := make([]int, n)
	size_two := make([]int, n)
	for i := 0; i < n; i++ {
		parent_one[i] = i
		parent_two[i] = i
		size_one[i] = 1
		size_two[i] = 1
	}
	for i := 0; i < n; i++ {
		union(nums[i], nums[nums[i]], &parent_one, &size_one)
		union(nums[i], nums[(nums[i]-1+n)%n], &parent_two, &size_two)
	}
	return min(score(parent_one, &size_one), score(parent_two, &size_two))
}

func score(parent []int, size *[]int) int {
	res := 0
	for i := 0; i < len(parent); i++ {
		if parent[i] == i && (*size)[i] > 1 {
			if find(0, &parent) != i { //i is not the parent of i
				(*size)[i]++
				res++
			}
			res += (*size)[i] - 1
		}
	}
	return res
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] > (*size)[y] {
		(*parent)[y] = x
		(*size)[x] += (*size)[y]
	} else {
		(*parent)[x] = y
		(*size)[y] += (*size)[x]
	}
}

func find(x int, parent *[]int) int {
	if x == (*parent)[x] {
		return x
	}
	(*parent)[x] = find((*parent)[x], parent)
	return (*parent)[x]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(sort_array_by_moving_to_empty_space([]int{4, 2, 0, 3, 1}))
}
