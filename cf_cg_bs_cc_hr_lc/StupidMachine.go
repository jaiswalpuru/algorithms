package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }

func main() {

	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {

		var n int
		scanf("%d\n", &n)

		var arr = make([][2]int, n)

		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i][0])
				break
			}
			scanf("%d ", &arr[i][0])
		}

		min := math.MaxInt64
		res := 0

		for i := 0; i < n; i++ {
			min = int(math.Min(float64(min), float64(arr[i][0])))
			res += min
		}

		fmt.Println(res)

	}

}
