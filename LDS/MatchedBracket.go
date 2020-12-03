package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdout)
	writer = bufio.NewWriter(os.Stdin)
)

func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }

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

	//some flags to keep track
	//f -> points to initial index of (
	//max depth md
	//
	var (
		ind, ind2             int
		l, ml, t, md, f, temp = 0, 0, 0, 0, 0, 0
	)

	for i := 0; i < n; i++ {
		if arr[i] == 1 { //(
			if l == 0 {
				f = i
			}
			l++
			temp++
			t++
		} else { // )
			l--
			if l != 0 {
				if temp > md {
					md = temp
					ind = i
				}
				t++
			} else {
				if t > ml {
					ml = t
					ind2 = f
					f = 0
				}
				t = 0
				if temp > md {
					md = temp
					ind = i
				}
			}
			temp = 0
		}

	}
	printf("%d %d %d %d\n", md, ind, ml+1, ind2+1)

}
