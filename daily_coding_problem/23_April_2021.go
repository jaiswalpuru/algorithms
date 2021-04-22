//Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.

package main

import (
	"fmt"
	"time"
)

func jobScheduler(f func(), n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Printf("%d milliseconds passed\n", n)
	f()
}

func main() {
	f := func() {
		fmt.Println("Called after 1000 milliseconds")
	}
	//calls f after 1000 millisecods
	jobScheduler(f, 1000)
}
