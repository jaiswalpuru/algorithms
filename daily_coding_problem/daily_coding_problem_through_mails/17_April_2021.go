// Given an array of integers, find the first missing positive integer in linear time and constant space.
//In other words, find the lowest positive integer that does not exist in the array.
//The array can contain duplicates and negative numbers as well.
// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

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

func main() {
	defer writer.Flush()

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
	//first segregate the positive integers and non positive integers
	j := 0
	for i := 0; i < n; i++ {
		if arr[i] <= 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	arr = arr[j:]
	//now mark all the postive numbers as visited by its negative representation
	n = n - j
	for i := 0; i < n; i++ {
		if int(math.Abs(float64(arr[i])))-1 < n && arr[int(math.Abs(float64(arr[i])))-1] > 0 {
			arr[int(math.Abs(float64(arr[i])))-1] *= -1
		}
	}
	fmt.Println(arr)
	i, f := 0, false
	for i = 0; i < n; i++ {
		if arr[i] > 0 {
			f = true
			break
		}
	}
	if f == true {
		fmt.Println(i + 1)
	} else {
		fmt.Println(n + 1)
	}
}
