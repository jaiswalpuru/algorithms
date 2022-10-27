package main

import "fmt"

//----------------bute force-------------------
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func shift_and_count(x, y int, a, b [][]int) int {
	left_shift_cnt, right_shift_cnt := 0, 0
	rb := 0 //row for matrix b
	for ra := y; ra < len(a); ra++ {
		cb := 0
		for ca := x; ca < len(a); ca++ {
			if a[ra][ca] == 1 && a[ra][ca] == b[rb][cb] {
				left_shift_cnt++
			}
			if a[ra][cb] == 1 && a[ra][cb] == b[rb][ca] {
				right_shift_cnt++
			}
			cb++
		}
		rb++
	}
	return max(left_shift_cnt, right_shift_cnt)
}

func image_overlap(img1, img2 [][]int) int {
	n := len(img1)
	max_overlap := 0
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			//move img1 up right and up left
			max_overlap = max(max_overlap, shift_and_count(x, y, img1, img2))
			//move img2 up right and up left
			max_overlap = max(max_overlap, shift_and_count(x, y, img2, img1))
		}
	}
	return max_overlap
}

//----------------bute force-------------------

//----------------linear transformation-------------------
func non_zero_cell(arr [][]int) [][2]int {
	res := [][2]int{}
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 1 {
				res = append(res, [2]int{i, j})
			}
		}
	}
	return res
}

func image_overlap_linear_transformation(a, b [][]int) int {
	a_cell := non_zero_cell(a)
	b_cell := non_zero_cell(b)
	mp := make(map[[2]int]int)
	res := 0
	for i := 0; i < len(a_cell); i++ {
		for j := 0; j < len(b_cell); j++ {
			t := [2]int{b_cell[j][0] - a_cell[i][0], b_cell[j][1] - a_cell[i][1]}
			mp[t]++
			res = max(res, mp[t])
		}
	}
	return res
}

//----------------linear transformation-------------------

func main() {
	fmt.Println(image_overlap_linear_transformation([][]int{
		{1, 1, 0}, {0, 1, 0}, {0, 1, 0},
	}, [][]int{
		{0, 0, 0}, {0, 1, 1}, {0, 0, 1},
	}))
}
