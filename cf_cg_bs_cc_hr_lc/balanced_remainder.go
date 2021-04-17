package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }
func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinElement(arr []int) int {
	m := int(math.MaxInt32)
	for i := 0; i < 3; i++ {
		m = min(m, arr[i])
	}
	return m
}

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for t > 0 {
		var n int
		scanf("%d\n", &n)
		arr := make([]int, n)
		cnt := make([]int, 3)
		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i])
			} else {
				scanf("%d ", &arr[i])
			}
			cnt[arr[i]%3]++ // storing the counts of remainders
		}

		ans := 0
		for getMinElement(cnt) != n/3 {
			for i := 0; i < 3; i++ {
				if cnt[i] > n/3 {
					ans++
					cnt[i]--
					cnt[(i+1)%3]++
				}
			}
		}
		printf("%d\n", ans)
		t--
	}
}
