package main

import "fmt"

const (
	A = "1"
	B = "2"
	C = "3"
)

func tower_of_hanoi(n int, a, b, c string) {
	if n >= 1 {
		tower_of_hanoi(n-1, a, c, b)
		fmt.Println("Move disk from ", a, c)
		tower_of_hanoi(n-1, b, a, c)
	}
}

func main() {
	tower_of_hanoi(3, A, B, C)
	fmt.Println("-------------------------------")
	tower_of_hanoi(4, A, B, C)
	fmt.Println("-------------------------------")
	tower_of_hanoi(5, A, B, C)
	fmt.Println("-------------------------------")
}
