package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {

	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {
		var n, k int
		scanf("%d %d\n", &n, &k)

		arr := make([]int, n)

		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i])
				break
			}
			scanf("%d ", &arr[i])
		}

		sumK, maxLen, t := 0, 1, 0
		mp := make(map[int]int)

		for i := 0; i < n; i++ {
			if mp[arr[i]] == 0 {
				sumK++
			}
			mp[arr[i]]++
			if sumK < k {
				maxLen = max(maxLen, i-t+1)
			}
			for sumK == k {
				mp[arr[t]]--
				if mp[arr[t]] == 0 {
					sumK--
				}
				t++
			}
		}
		printf("%d\n", maxLen)
	}
}
