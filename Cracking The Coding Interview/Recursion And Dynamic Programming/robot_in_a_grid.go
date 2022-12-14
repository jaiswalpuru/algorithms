package main

import (
	"fmt"
	"strconv"
)

//------------------brute force--------------------------
func robot_in_a_grid(arr [][]int) [][]int {
	paths := [][]int{}
	if recur(arr, len(arr)-1, len(arr[0])-1, &paths) {
		return paths
	}
	return nil
}

func recur(arr [][]int, i, j int, paths *[][]int) bool {
	if i < 0 || j < 0 || arr[i][j] == 1 {
		return false
	}
	if recur(arr, i-1, j, paths) || recur(arr, i, j-1, paths) || (i == 0 && j == 0) {
		*paths = append(*paths, []int{i, j})
		return true
	}
	return false
}

//------------------brute force--------------------------

//------------------efficient approach--------------------------
func robot_in_a_grid_eff(arr [][]int) [][]int {
	false_path := make(map[string]bool)
	paths := [][]int{}
	if _memo(arr, len(arr)-1, len(arr[0])-1, &paths, &false_path) {
		return paths
	}
	return nil
}

func _memo(arr [][]int, i, j int, paths *[][]int, false_path *map[string]bool) bool {
	if i < 0 || j < 0 || arr[i][j] == 1 {
		return false
	}
	if (*false_path)[to_string(i, j)] {
		return false
	}
	if (i == 0 && j == 0) || _memo(arr, i-1, j, paths, false_path) || _memo(arr, i, j-1, paths, false_path) {
		*paths = append(*paths, []int{i, j})
		return true
	}
	(*false_path)[to_string(i, j)] = true
	return false
}

func to_string(a, b int) string {
	s1, s2 := "", ""
	if a == 0 {
		s1 = "nil"
	} else {
		s1 = strconv.Itoa(a)
	}
	if b == 0 {
		s2 = "nil"
	} else {
		s2 = strconv.Itoa(b)
	}
	return s1 + "->" + s2
}

//------------------efficient approach--------------------------

func main() {
	fmt.Println(robot_in_a_grid_eff([][]int{
		{0, 0, 1, 0},
		{1, 0, 0, 0},
	}))
}
