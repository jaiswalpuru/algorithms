/*
An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting
from the pixel image[sr][sc].

To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the
starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those
pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with newColor.

Return the modified image after performing the flood fill.

Example 1:
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel),
all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are
colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

Example 2:
Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
Output: [[2,2,2],[2,2,2]]
*/

package main

import "fmt"

func flood_fill(image [][]int, sr, sc, new_color int) [][]int {
	formed_color := image[sr][sc]

	if formed_color == new_color {
		return image
	}
	image[sr][sc] = new_color
	if has_same_color(image, sr+1, sc, formed_color) {
		flood_fill(image, sr+1, sc, new_color)
	}
	if has_same_color(image, sr-1, sc, formed_color) {
		flood_fill(image, sr-1, sc, new_color)
	}
	if has_same_color(image, sr, sc+1, formed_color) {
		flood_fill(image, sr, sc+1, new_color)
	}
	if has_same_color(image, sr, sc-1, formed_color) {
		flood_fill(image, sr, sc-1, new_color)
	}
	return image
}

func has_same_color(image [][]int, i, j, formed_color int) bool {
	return i >= 0 && i < len(image) && j >= 0 && j < len(image[i]) && image[i][j] == formed_color
}

func main() {
	fmt.Println(flood_fill([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2))
}
