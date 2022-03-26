package main

import "fmt"

//recursive approach
func build_seg_tree(arr []int, tree_index int, lo, hi int, tree *[]int) {
	if lo == hi {
		(*tree)[tree_index] = arr[lo]
		return
	}

	mid := (lo+hi)/2
	build_seg_tree(arr, 2*tree_index+1, lo, mid, tree)
	build_seg_tree(arr, 2*tree_index+2, mid+1, hi, tree)
	(*tree)[tree_index] = (*tree)[2*tree_index+1] + (*tree)[2*tree_index+2]
}

//range/query on an interval or segment of data
func query_seg_tree(tree_index int, lo, hi, i, j int, tree []int ) int {
	if lo > j || hi < i{
		return 0
	}

	if i<=lo && j >= hi {
		return tree[tree_index]
	}

	mid := (lo+hi) / 2 //partial overlap of current segment and queried range. Recurse deeper

	if i > mid {
		return query_seg_tree(2*tree_index+2, mid+1, hi, i, j, tree)
	} else if j <= mid {
		return query_seg_tree(2*tree_index+1, lo, mid, i, j,tree)
	}

	left_query := query_seg_tree(2*tree_index+1, lo, mid, i, mid, tree)
	right_query := query_seg_tree(2*tree_index+2, mid+1, hi, mid+1, j, tree)

	return left_query+right_query
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
	tree := make([]int, 2*len(a))
	build_seg_tree(a, 0, 0,len(a)-1, &tree)
	fmt.Println(tree)
}