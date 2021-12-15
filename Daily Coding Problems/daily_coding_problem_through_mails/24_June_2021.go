/*
	Good morning! Here's your coding interview problem for today.

	This problem was asked by Two Sigma.

	Using a function rand7() that returns an integer from 1 to 7 (inclusive) with uniform probability,
	implement a function rand5() that returns an integer from 1 to 5 (inclusive).
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand7() int {
	return 1 + rand.Intn(7)
}

func rand5() int {
	val := 5*rand7() + rand7() - 5
	if val <= 49 {
		return val%5 + 1
	}
	return rand5()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Print(rand5(), " ")
	}
	fmt.Println()
}
