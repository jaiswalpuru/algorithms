package main

import "fmt"

func fibonacci_number(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func main() {
	fmt.Println(fibonacci_number(30))
}
