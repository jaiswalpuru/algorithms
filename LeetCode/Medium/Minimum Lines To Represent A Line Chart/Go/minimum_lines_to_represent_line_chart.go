package main

import "fmt"
import "sort"

type I [][]int

func (a I) Len() int           { return len(a) }
func (a I) Less(i, j int) bool { return a[i][0] < a[j][0] }
func (a I) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func minimum_lines(prices [][]int) int {
	var p I
	p = prices
	sort.Sort(p)

	if p.Len() == 1 {
		return 0
	}
	n := p.Len()
	lines := 1
	for i := 2; i < n; i++ {
		x1, x2, x3 := p[i][0], p[i-1][0], p[i-2][0]
		y1, y2, y3 := p[i][1], p[i-1][1], p[i-2][1]

		v1, v2 := (y3-y2)*(x2-x1), (y2-y1)*(x3-x2)
		if v1 != v2 {
			lines++
		}
	}
	return lines
}

func main() {
	fmt.Println(minimum_lines([][]int{
		{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1},
	}))
}
