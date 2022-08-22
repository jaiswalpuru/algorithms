package main

import (
	"fmt"
	"math"
)

func power_of_four(n int) bool {
	val := math.Log2(float64(n)) / 2
	return val-float64(int(val)) == 0
}

func main() {
	fmt.Println(power_of_four(16))
}
