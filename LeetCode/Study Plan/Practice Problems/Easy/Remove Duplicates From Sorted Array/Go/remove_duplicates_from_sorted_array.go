package main

import "fmt" 

func remove_duplicates_from_sorted_array(arr []int) []int {
	i,j := 0,0
	n := len(arr)
	for i<n && j<n  {
		curr := arr[j]
		for j<n && curr == arr[j]{
			j++
		}
		arr[i] = arr[j-1]
		i++
	}
	return i
}

func main () {
	fmt.Println(remove_duplicates_from_sorted_array([]int{1,1,2}))
}