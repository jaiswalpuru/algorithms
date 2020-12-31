package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inversion = 0
	reader    = bufio.NewReader(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	scanf     = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf    = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

func merge(left, right []int) []int {
	result := []int{}
	a, b := 0, 0
	for (a < len(left)) && (b < len(right)) {
		if left[a] <= right[b] {
			result = append(result, left[a])
			a++
		} else {
			inversion += len(left) - a
			result = append(result, right[b])
			b++
		}
	}
	for a < len(left) {
		result = append(result, left[a])
		a++
	}
	for b < len(right) {
		result = append(result, right[b])
		b++
	}
	return result
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	middle := len(a) / 2
	left := mergeSort(a[:middle])
	right := mergeSort(a[middle:])
	return merge(left, right)
}

func main() {
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &arr[i])
			break
		}
		scanf("%d ", &arr[i])
	}

	_ = mergeSort(arr)
	fmt.Println(inversion)
}
