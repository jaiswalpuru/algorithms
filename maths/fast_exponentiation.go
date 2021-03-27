package main

import "fmt"

func pow(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		halfAns := pow(a, n/2)
		return halfAns * halfAns
	} else {
		halfAns := pow(a, n/2) //floor value of n/2
		return halfAns * halfAns * a
	}
}

func main() {
	fmt.Println(pow(2, 13))
}
