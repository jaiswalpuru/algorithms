/*
	Using a function rand5() that returns an integer from 1 to 5 (inclusive) with uniform probability,
	implement a function rand7() that returns an integer from 1 to 7 (inclusive).
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//given method rand5() returns random number from [1,5] implemented for testing purpose
func rand5() int {
	rand.Seed(time.Now().Local().UnixNano())
	return 1 + rand.Intn(5)
}

func rand7() int {
	val := 5*rand5() + rand5() - 5
	if val < 22 {
		return val%7 + 1
	}
	return rand7()
}

func main() {
	fmt.Println(rand7())
}
