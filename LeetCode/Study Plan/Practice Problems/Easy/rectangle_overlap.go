/*
An axis-aligned rectangle is represented as a list [x1, y1, x2, y2], where (x1, y1) is the coordinate of its
bottom-left corner, and (x2, y2) is the coordinate of its top-right corner. Its top and bottom edges are parallel
to the X-axis, and its left and right edges are parallel to the Y-axis.

Two rectangles overlap if the area of their intersection is positive. To be clear, two rectangles that only touch at
the corner or edges do not overlap.

Given two axis-aligned rectangles rec1 and rec2, return true if they overlap, otherwise return false.
*/

package main

import "fmt"

func is_overlap(r1 []int, r2 []int) bool {
	r1x1, r1y1, r1x2, r1y2 := r1[0], r1[1], r1[2], r1[3]
	r2x1, r2y1, r2x2, r2y2 := r2[0], r2[1], r2[2], r2[3]
	if r1x1 == r1x2 || r1y1 == r1y2 || r2x1 == r2x2 || r2y1 == r2y2 {
		return false
	}

	return !(r1x2 <= r2x1 || r1y2 <= r2y1 || r1x1 >= r2x2 || r1y1 >= r2y2)

}

func main() {
	fmt.Println(is_overlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
}
