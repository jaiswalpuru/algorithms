package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {
		var n, k int
		scanf("%d %d\n", &n, &k)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i])
				break
			}
			scanf("%d ", &arr[i])
		}

		ans := 0
		for i := 0; i < n; i++ {
			freq := make([]int, 2001)
			for k := 0; k < 2001; k++ {
				freq[k] = 0
			}
			temp := []int{}
			for j := i; j < n; j++ {
				m := int(math.Ceil(float64(k) / float64(j-i+1))) //calculating the value of m (j-i+1) is the length of the subarray
				k1 := int(math.Ceil(float64(k) / float64(m)))
				k1-- //due to zero based indexing
				//now find the kth smallest number in O(1) or O(logn)
				freq[arr[j]]++
				temp = append(temp, arr[j]) //complexity increases for larger inputs future implementation on ordered sets
				sort.Ints(temp)
				element := temp[k1]
				frequency := freq[element]
				if freq[frequency] > 0 {
					ans++
				}
			}
		}
		printf("%d\n", ans)
	}
}
