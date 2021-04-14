// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var scanner *bufio.Scanner = bufio.NewScanner(reader)

func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }
func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }

func main() {
	defer writer.Flush()

	var n, k int
	scanf("%d %d\n", &n, &k)

	arr := make([]int, n)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &arr[i])
		} else {
			scanf("%d ", &arr[i])
		}
		mp[arr[i]] = arr[i]
	}

	for v := range mp {
		if _, ok := mp[k-v]; ok {
			printf("Pair {%d,%d}\n", v, k-v)
			break
		}
	}
}
