package main

import ("fmt"; "sort")

func boats_to_save_people(arr []int, limit int) int {
	sort.Ints(arr)
	n := len(arr)
	boat_cnt := 0

	i,j := 0, n-1

	for i<=j {
		boat_cnt++
		if arr[i]+arr[j] <= limit {
			i++
		}
		j--
	}	

	return boat_cnt
}

func main() {
	fmt.Println(boats_to_save_people([]int{3,2,2,1},3))
}