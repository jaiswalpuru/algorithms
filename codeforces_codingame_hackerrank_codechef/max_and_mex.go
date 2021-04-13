package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type i64 int64

func max(a, b i64) i64 {
	if a > b {
		return a
	}
	return b
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)
	for t > 0 {
		var n, k, i i64
		scanf("%d %d\n", &n, &k)
		arr := make([]i64, n)
		b := i64(math.MinInt64)
		mp := make(map[i64]bool)
		for i = 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i])
			} else {
				scanf("%d ", &arr[i])
			}
			mp[arr[i]] = true
			b = max(b, arr[i])
		}
		if k == 0 {
			fmt.Println(n)
		} else {
			a := i64(0) //mex
			for i = 0; i <= n; i++ {
				if mp[i] == false {
					a = i
					break
				}
			}
			if a > b {
				fmt.Println(n + k)
			} else {
				ans := len(arr)
				val := (a + b + 1) / 2
				if mp[val] == false {
					ans++
				}
				fmt.Println(ans)
			}
		}
		t--
	}
}
