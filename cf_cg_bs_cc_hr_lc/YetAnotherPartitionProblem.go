package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

func main() {
	var n, q, t, ind, val int
	scanf("%d %d\n", &n, &q)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &arr[i])
			break
		}
		scanf("%d ", &arr[i])
	}

	set := []int{}
	mp := make(map[int]int)
	if _, ok := mp[0]; !ok {
		set = append(set, 0)
	}
	for i := 1; i < n; i++ {
		if arr[i]%arr[i-1] != 0 {
			if _, ok := mp[i]; !ok {
				set = append(set, i)
			}
		}
	}

	sort.Ints(set)

	for i := 0; i < q; i++ {
		scanf("%t ", &t)
		if t == 1 {
			scanf("%d %d\n", &ind, &val)

		} else {
			scanf("%d\n", &ind)
			sort.Search(len(set))
		}
	}
}
