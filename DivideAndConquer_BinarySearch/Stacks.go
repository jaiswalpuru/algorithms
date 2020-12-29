package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

type p []int

//Sort not required in this problem
func (a p) Len() int           { return len(a) }
func (a p) Less(i, j int) bool { return a[i] < a[j] }
func (a p) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func strictlyLowerBoundBSearch(a []int, val, l, h int) int {
	if len(a) == 0 {
		return 0
	}
	if val < a[l] {
		return l
	} else if val >= a[h] {
		return h + 1
	} else {
		mid := (l + h) / 2
		if a[mid] <= val {
			l = mid + 1
		} else {
			h = mid - 1
		}
		return strictlyLowerBoundBSearch(a, val, l, h)
	}
}

func main() {
	defer writer.Flush()
	//the disk in the single stack must be ordered by their radiuses in a strictly increasing manner
	//such that the top most diks has the smallest radius
	//If there is at least one stack such that Mike can put the current disk on the top of the stack without making
	//it invalid, then he chooses the stack with the smallest top disk radius strictly greater than the radius of
	//the current disk, and puts the current disk on top of that stack.
	var (
		t, n int
		k    = 0
	)
	scanf("%d\n", &t)

	for ; t > 0; t-- {
		scanf("%d\n", &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i])
				break
			}
			scanf("%d ", &arr[i])
		}

		temp := []int{}
		for i := 0; i < n; i++ {
			if len(temp) > 0 {
				k = strictlyLowerBoundBSearch(temp, arr[i], 0, len(temp)-1)
				//fmt.Printf("i=%d, K=%d, temp=%+v, arr[%d] = %d\n", i, k, temp, i, arr[i]) - debug
				if k >= len(temp) {
					temp = append(temp, arr[i])
				} else {
					temp[k] = arr[i]
				}
			} else {
				temp = append(temp, arr[i])
			}
		}
		printf("%d", len(temp))
		for j := 0; j < len(temp); j++ {
			if j == len(temp)-1 {
				printf(" %d\n", temp[j])
				break
			}
			printf(" %d", temp[j])
		}
	}
}
