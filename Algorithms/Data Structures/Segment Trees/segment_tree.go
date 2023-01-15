package main

import "fmt"

//recursive approach
func build_seg_tree(arr []int, tree_index int, low, high int, tree *[]int) {
	if low== high {
		(*tree)[tree_index] = arr[low]
		return
	}

	mid := (low+high)/2
	build_seg_tree(arr, 2*tree_index+1, low, mid, tree)
	build_seg_tree(arr, 2*tree_index+2, mid+1, high, tree)
	(*tree)[tree_index] = (*tree)[2*tree_index+1] + (*tree) [2*tree_index+2]
}

//range/query on an interval or segment of data
func query_seg_tree(ind, low, high, l, r int, tree []int) int {
	if low > l && high < r { //node completly lies inside this range
		return tree[ind]
	}

	if high < l || low > r {
		return 0
	}

	mid := (low + high) /2
	left := query_seg_tree(2*ind+1, low, mid, l, r, tree)
	right := query_seg_tree(2*ind+2, mid+1, high, l, r, tree)

	return left+right
}

//update value of an element
func update_value(tree_index int, lo, hi, index, val int, tree *[]int) {
	if lo == hi {
		(*tree)[tree_index] = val
		return
	}

	mid := (lo + hi) / 2
	if index > mid {
		update_value(2*tree_index+2, mid+1, hi, index, val, tree)
	}else if index <= mid {
		update_value(2*tree_index+1, lo, mid, index, val, tree)
	}

	(*tree)[tree_index] = (*tree)[2*tree_index+1] + (*tree)[2*tree_index+2]
}

func main() {
	a := []int{1,3,5}
	tree := make([]int, 4*len(a))
	build_seg_tree(a, 0, 0,len(a)-1, &tree)
	fmt.Println(tree)
}