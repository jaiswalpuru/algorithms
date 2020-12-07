package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type conDetails struct {
	start, end int
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

	defer writer.Flush()

	var (
		n, x, y int
	)

	scanf("%d %d %d\n", &n, &x, &y)

	contest := make([]conDetails, n)
	v := make([]int, x)
	w := make([]int, y)

	for i := 0; i < n; i++ {
		scanf("%d %d\n", &contest[i].start, &contest[i].end)
	}

	for i := 0; i < x; i++ {
		if i == x-1 {
			scanf("%d\n", &v[i])
			break
		}
		scanf("%d ", &v[i])
	}

	for i := 0; i < y; i++ {
		if i == y-1 {
			scanf("%d\n", &w[i])
			break
		}
		scanf("%d ", &w[i])
	}

	sort.Ints(v)
	sort.Ints(w)

	ans := math.MaxInt64

	for i := 0; i < n; i++ {
		t1 := sort.SearchInts(v, contest[i].start)
		t2 := sort.SearchInts(w, contest[i].end)
		//println(contest[i].start, contest[i].end, t1, t2)
		if t1 >= len(v) || v[t1] != contest[i].start {
			t1--
		}
		if t2 >= len(w) {
			t2--
		} else if w[t2] != contest[i].end {
			t2++
		}
		ans = min(ans, w[t2]-v[t1]+1)
	}

	printf("%d\n", ans)
}
