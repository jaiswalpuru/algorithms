package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

func main() {
	defer writer.Flush()

	var (
		t, n, m, k, r, c int
		//dx and dy are present to check if plants are present in adjacent coordinates
		dx  = []int{1, 0, -1, 0}
		dy  = []int{0, 1, 0, -1}
		ans = 0
	)

	scanf("%d\n", &t)
	for ; t > 0; t-- {
		ans = 0
		mp := make(map[int]map[int]int)
		scanf("%d %d %d\n", &n, &m, &k)
		for ; k > 0; k-- {
			scanf("%d %d\n", &r, &c)
			for i := 0; i < 4; i++ {
				nx := r + dx[i]
				ny := c + dy[i]
				if mp[nx][ny] == 1 {
					ans += -1
				} else {
					ans++
				}
			}
			if _, ok := mp[r]; ok {
				mp[r][c] = 1
			} else {
				mp[r] = make(map[int]int)
				mp[r][c] = 1
			}
		}
		printf("%d\n", ans)
	}
}
