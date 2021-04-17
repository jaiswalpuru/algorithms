package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	mod    = 1000000007
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
)

type rankInd struct {
	rank int
	ind  int
}

func main() {
	defer writer.Flush()

	var (
		n, k int
		t    = -1
	)

	scanf("%d %d\n", &n, &k)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &arr[i])
			break
		}
		scanf("%d ", &arr[i])
	}

	f := 1
	temp := make([]rankInd, len(arr))

	t++
	temp[t].ind = 0
	temp[t].rank = arr[0]

	for i := 1; i < n; i++ {
		for t >= 0 {
			if temp[t].rank > arr[i] {
				f *= (i - temp[t].ind + 1)
				f %= mod
				t--
			} else {
				break
			}
		}
		t++
		temp[t].ind = i
		temp[t].rank = arr[i]
	}

	printf("%d\n", f%mod)

}
