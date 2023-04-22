package main

import "fmt"

func check_if_it_is_a_straight_line(coordinates [][]int) bool {
	c1, c2 := coordinates[0], coordinates[1]
	for i := 2; i < len(coordinates); i++ {
		if (coordinates[i][1]-c2[1])*(c2[0]-c1[0]) != (coordinates[i][0]-c2[0])*(c2[1]-c1[1]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(check_if_it_is_a_straight_line([][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7},
	}))
}
