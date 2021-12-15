/*
	Assume you have access to a function toss_biased() which returns 0 or 1 with a
	probability that's not 50-50 (but also not 0-100 or 100-0). You do not know the bias of the coin.

	Write a function to simulate an unbiased coin toss.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TAILS = 0
	HEADS = 1
)

func biased() int {
	val := rand.Int() % 100

	if val <= 79 {
		return HEADS
	}
	return TAILS
}

func generate() int {
	for {
		first := biased()
		second := biased()
		if first != second {
			return first
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	x, y := 0, 0
	for i := 0; i < 100000; i++ {
		val := generate()
		if val == 0 {
			x++
		} else {
			y++
		}
	}
	fmt.Println(float64(x)/1000.0, float64(y)/1000.0)
}
