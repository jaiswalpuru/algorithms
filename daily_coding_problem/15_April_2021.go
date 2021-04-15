/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
*/

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

func initArr(arr *[]int, val int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = val
	}
}

func main() {
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	//using division method
	arr := make([]int, n)
	arrAlt := make([]int, n)
	finNum := 1
	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &arr[i])
		} else {
			scanf("%d ", &arr[i])
		}
		arrAlt[i] = arr[i]
		finNum *= arr[i]
	}

	for i := 0; i < n; i++ {
		arr[i] = finNum / arr[i]
	}
	fmt.Println(arr)

	//without division
	temp := 1
	prod := make([]int, n)
	initArr(&prod, 1)
	for i := 0; i < n; i++ {
		prod[i] = temp
		temp *= arrAlt[i]
	}
	temp = 1
	for i := n - 1; i >= 0; i-- {
		prod[i] *= temp
		temp *= arrAlt[i]
	}
	fmt.Println(prod)
}
