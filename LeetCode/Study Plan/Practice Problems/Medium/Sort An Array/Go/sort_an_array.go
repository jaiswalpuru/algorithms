package main

import "fmt"

//---------------------------- applying merge sort ----------------------------
func sort_an_array(arr []int) []int {
	if len(arr)<=1 {
		return arr
	}
	pivot := len(arr)/2
	left := sort_an_array(arr[:pivot])
	right := sort_an_array(arr[pivot:])
	return merge(left, right)
}

func merge(a, b []int) []int{
	res := make([]int, len(a)+len(b))
	l,r,p := 0,0,0
	for l < len(a) && r < len(b) {
		if a[l] < b[r] {
			res[p] = a[l]
			l++
		} else {
			res[p] = b[r]
			r++
		}
		p++
	}

	for l < len(a) {
		res[p] = a[l]
		l++
		p++
	}

	for r < len(b) {
		res[p] = b[r]
		r++
		p++
	}

	return res
}
//-----------------------------------------------------------------------------

func main() {
	fmt.Println(sort_an_array([]int{5,2,3,1}))
}