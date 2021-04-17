package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	val, ind int
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

type xSorter []pair

func (a xSorter) Len() int           { return len(a) }
func (a xSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a xSorter) Less(i, j int) bool { return a[i].val < a[j].val }

func main() {
	defer writer.Flush()

	var (
		n, m int
	)

	scanf("%d %d\n", &n, &m)
	a := make([]pair, n)
	b := make([]pair, m)

	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &a[i].val)
			a[i].ind = i
			break
		}
		scanf("%d ", &a[i].val)
		a[i].ind = i
	}

	for i := 0; i < m; i++ {
		if i == m-1 {
			scanf("%d\n", &b[i].val)
			b[i].ind = i
			break
		}
		scanf("%d ", &b[i].val)
		b[i].ind = i
	}

	sort.Sort(xSorter(a))
	sort.Sort(xSorter(b))

	fmt.Println(a, b)

	for i := 0; i < m; i++ {
		printf("%d %d\n", a[0].ind, b[i].ind)
	}

	for i := 1; i < n; i++ {
		printf("%d %d\n", a[i].ind, b[m-1].ind)
	}

	// mp := make(map[int]int)
	// i, j, f := 0, len(a)-1, false

	// for len(mp) < n+m-1 {
	// 	f = !f
	// 	var tmp pair
	// 	if f {
	// 		tmp.val = a[i].val
	// 		tmp.ind = a[i].ind
	// 	} else {
	// 		tmp.val = a[j].val
	// 		tmp.ind = a[j].ind
	// 	}
	// 	for v := range b {
	// 		sum := tmp.val + b[v].val
	// 		if _, ok := mp[sum]; !ok {
	// 			printf("%d %d\n", tmp.ind, b[v].ind)
	// 			mp[sum] = 1
	// 		}
	// 		if len(mp) == n+m-1 {
	// 			break
	// 		}
	// 	}

	// 	if f {
	// 		i++
	// 	} else {
	// 		j--
	// 	}
	// }

}
