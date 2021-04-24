//Given a list of integers, write a function that returns the largest sum of non-adjacent numbers.
//Numbers can be 0 or negative.
//For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5.
//[5, 1, 1, 5] should return 10, since we pick 5 and 5.

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }
func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }

func main() {
	writer.Flush()

	var n int
	scanf("%d\n", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &arr[i])
		} else {
			scanf("%d ", &arr[i])
		}
	}

	incl, excl := arr[0], 0
	// incl tells what can be the maximum sum including that number and excl tells what maximum sum that can be done excluding that number
	for i := 1; i < n; i++ {
		temp := incl
		incl = max(incl, arr[i]+excl)
		excl = temp
	}
	fmt.Println(incl)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
