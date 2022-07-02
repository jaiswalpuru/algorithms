package main

import (
	"fmt"
	"math"
	"sort"
)

var mod = int(1e9 + 7)

func maximum_area_of_piece_of_cake_after_horizontal_and_vertical_cuts(h, w int, horizontal_cut []int, vertical_cut []int) int {
	sort.Ints(horizontal_cut)
	sort.Ints(vertical_cut)
	max_height := int(math.Max(float64(horizontal_cut[0]), float64(h-horizontal_cut[len(horizontal_cut)-1])))
	max_width := int(math.Max(float64(vertical_cut[0]), float64(w-vertical_cut[len(vertical_cut)-1])))
	for i := 0; i < len(horizontal_cut)-1; i++ {
		max_height = int(math.Max(float64(max_height), float64(horizontal_cut[i+1]-horizontal_cut[i])))
	}
	for i := 0; i < len(vertical_cut)-1; i++ {
		max_width = int(math.Max(float64(max_width), float64(vertical_cut[i+1]-vertical_cut[i])))
	}
	return (max_width % mod * max_height % mod) % mod
}

func main() {
	fmt.Println(maximum_area_of_piece_of_cake_after_horizontal_and_vertical_cuts(5, 4, []int{1, 2, 4}, []int{1, 3}))
}
