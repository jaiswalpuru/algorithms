package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

func fillArray(a []int, n int) (sum int) {
	sum = 0
	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &a[i])
			sum += a[i]
			break
		}
		scanf("%d ", &a[i])
		sum += a[i]
	}
	return
}

func main() {
	defer writer.Flush()

	var (
		t, n, k int
		ht      []int
	)
	scanf("%d\n", &t)

	for ; t > 0; t-- {
		ans, j := 0, 0
		scanf("%d %d\n", &n, &k)
		ht = make([]int, n)
		sum := fillArray(ht, n)
		if sum < 2*k || n == 1 {
			printf("%d\n", -1)
		} else {
			sort.Ints(ht)
			if ht[n-1] >= k && ht[n-2] >= k {
				printf("%d\n", 2)
			} else if ht[n-1] >= k && ht[n-2] < k {
				sum = 0
				ans++
				for i := n - 2; i >= 0; i-- {
					sum += ht[i]
					ans++
					if sum >= k {
						break
					}
				}
				printf("%d\n", ans)
			} else if ht[n-1] < k && ht[n-2] < k {
				sum = 0
				for j = n - 1; j >= 0; j-- {
					sum += ht[j]
					ans++
					if sum >= 2*k {
						break
					}
				}
				// m := 0
				// for i := j; i < n; i++ {
				// 	if (sum - ht[i]) >= 2*k {
				// 		ans--
				// 		sum -= ht[i]
				// 	} else if (sum-ht[i]) >= 2*k && (sum+ht[m]) >= 2*k {

				// 	}
				// }
				printf("%d\n", ans)
			}
		}
	}
}
