package main

//explanation given by codechef on : youtube.com/watch?v=7Q2Nlf-fT4o

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pos struct {
	x, y int
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
)

type xSorter []pos

func (a xSorter) Len() int           { return len(a) }
func (a xSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a xSorter) Less(i, j int) bool { return a[i].x < a[j].x }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	coord := make([]pos, n)

	for i := 0; i < n; i++ {
		scanf("%d %d\n", &coord[i].x, &coord[i].y)
	}

	sort.Sort(xSorter(coord))

	t := -1
	stack := make([]pos, len(coord))
	rb := make([]int, n)
	lb := make([]int, n)

	//found the left bound for each point where i is pivot
	for i := 0; i < n; i++ {
		for t != -1 && stack[t].x > coord[i].y {
			p := &pos{x: stack[t].x, y: stack[t].y}
			t--
			rb[p.y] = coord[i].x
		}
		t++
		stack[t] = pos{x: coord[i].y, y: i}
	}

	for t != -1 {
		p := pos{x: stack[t].x, y: stack[t].y}
		t--
		rb[p.y] = 100000
	}

	//found the right bound for each point where i is pivot
	for i := n - 1; i >= 0; i-- {
		for t != -1 && stack[t].x > coord[i].y {
			p := pos{x: stack[t].x, y: stack[t].y}
			t--
			lb[p.y] = coord[i].x
		}
		t++
		stack[t] = pos{x: coord[i].y, y: i}
	}

	for t != -1 {
		p := pos{x: stack[t].x, y: stack[t].y}
		t--
		lb[p.y] = 0
	}

	//using the left and right bound using the left and right of the rectangle, we found max area
	var ans = 0
	for i := 0; i < n; i++ {
		left := lb[i]
		right := rb[i]
		height := coord[i].y
		width := right - left
		ans = max(ans, height*width)
	}
	//edge case where top of rectangle is plane itself
	ans = max(ans, 500*coord[0].x)
	for i := 0; i < n-1; i++ {
		dist := coord[i+1].x - coord[i].x // distance between consecutive elements
		ans = max(ans, 500*dist)          //the top as the edge of the plane (height 500) and the width as the distance between the points
	}
	ans = max(ans, (100000-coord[n-1].x)*500) //distance between the right edge of the plane and the last point
	printf("%d\n", ans)
}
