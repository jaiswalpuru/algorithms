package main

import "fmt"

func rectangle_area(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	x1 := max(ax1, bx1)
	y1 := max(ay1, by1)
	x2 := min(ax2, bx2)
	y2 := min(ay2, by2)
	area_a1 := area(ax1, ay1, ax2, ay2)
	area_a2 := area(bx1, by1, bx2, by2)
	area_a3 := 0
	if x2 > x1 && y2 > y1 {
		area_a3 = area(x1, y1, x2, y2)
	}
	return area_a1 + area_a2 - area_a3
}

func area(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(rectangle_area(-3, 0, 3, 4, 0, -1, 9, 2))
}
