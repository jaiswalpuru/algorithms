/*
Given an array points where points[i] = [xi, yi] represents a point on the X-Y plane, return true if these points are a
boomerang.

A boomerang is a set of three points that are all distinct and not in a straight line.
*/

package main

import "fmt"

//Note: Three points are said to be collinear if the area of triangle formed by the three points == 0
//Area of triangle = 1/2 ((x1-x2)(y2-y3)-(x2-x3)(y1-y2))

func valid_boomerang(arr [][]int) bool {
	x1, x2, x3 := arr[0][0], arr[1][0], arr[2][0]
	y1, y2, y3 := arr[0][1], arr[1][1], arr[2][1]

	return ((x1-x2)*(y2-y3) - (x2-x3)*(y1-y2)) != 0
}

func main() {
	fmt.Println([][]int{{1, 1}, {2, 2}, {3, 3}})
}
