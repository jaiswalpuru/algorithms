package main

import "fmt"

func maximum_points_on_a_line(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	res := 2
	for i := 0; i < len(points); i++ {
		slopes := make(map[float64]map[[2]int]bool)
		slope_horizontal_line := make(map[int]map[[2]int]bool)
		for j := i + 1; j < len(points); j++ {
			if points[j][0]-points[i][0] != 0 {
				slope := get_slope(points[i][0], points[j][0], points[i][1], points[j][1])
				if _, ok := slopes[slope]; !ok {
					slopes[slope] = make(map[[2]int]bool)
					slopes[slope][[2]int{points[i][0], points[i][1]}] = true
				}
				slopes[slope][[2]int{points[j][0], points[j][1]}] = true
				res = max(res, len(slopes[slope]))
			} else {
				if _, ok := slope_horizontal_line[points[i][0]]; !ok {
					slope_horizontal_line[points[i][0]] = make(map[[2]int]bool)
					slope_horizontal_line[points[i][0]][[2]int{points[i][0], points[i][1]}] = true
				}
				slope_horizontal_line[points[i][0]][[2]int{points[j][0], points[j][1]}] = true
				res = max(res, len(slope_horizontal_line[points[i][0]]))
			}
		}
	}
	return res
}

func get_slope(x1, x2, y1, y2 int) float64 {
	return float64(y2-y1) / float64(x2-x1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximum_points_on_a_line([][]int{
		{1, 1}, {2, 2}, {3, 3},
	}))
}
