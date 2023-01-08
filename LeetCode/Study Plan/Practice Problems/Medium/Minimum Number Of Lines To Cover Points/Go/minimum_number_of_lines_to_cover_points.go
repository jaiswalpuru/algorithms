package main

import (
	"fmt"
	"math"
)

func minimum_number_of_lines_to_cover_points(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	return backtrack([]Pair{}, points)
}

type Pair struct {
	point []int
	slope float64
}

func backtrack(lines []Pair, points [][]int) int {
	if len(points) == 0 {
		return len(lines)
	}
	curr_point := points[0]
	for i := 0; i < len(lines); i++ {
		s := slope(curr_point, lines[i].point)
		if s == lines[i].slope {
			return backtrack(lines, points[1:])
		}
	}
	if len(points) == 1 {
		return len(lines) + 1
	}
	res := math.MaxInt64
	for i := 1; i < len(points); i++ {
		curr := points[i]
		s := slope(curr_point, curr)
		lines = append(lines, Pair{curr, s})
		_point := [][]int{}
		_point = append(_point, points[1:i]...)
		_point = append(_point, points[i+1:]...)
		res = min(res, backtrack(lines, _point))
		lines = lines[:len(lines)-1]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func slope(p1, p2 []int) float64 {
	if p1[0] == p2[0] {
		return math.Inf(1)
	}
	return float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
}

func main() {
	fmt.Println(minimum_number_of_lines_to_cover_points([][]int{
		{0, 1}, {2, 3}, {4, 5}, {4, 3},
	}))
}
