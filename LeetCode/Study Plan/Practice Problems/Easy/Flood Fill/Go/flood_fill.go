package main

import "fmt"

func flood_fill(arr [][]int, sr, sc int, new_color int) [][]int {
	_flood_fill(&arr, sr, sc, arr[sr][sc], new_color)
	return arr
}

func _flood_fill(arr *[][]int, i, j int, color, new_color int) {
	if i >= len(*arr) || j >= len((*arr)[0]) || i < 0 || j < 0 || (*arr)[i][j] != color {
		return
	}
	(*arr)[i][j] = new_color
	_flood_fill(arr, i-1, j, color, new_color)
	_flood_fill(arr, i, j-1, color, new_color)
	_flood_fill(arr, i+1, j, color, new_color)
	_flood_fill(arr, i, j+1, color, new_color)
}

func main() {
	fmt.Println(flood_fill([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2))
}
