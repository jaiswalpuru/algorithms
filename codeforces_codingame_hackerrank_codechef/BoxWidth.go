package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var scanner *bufio.Scanner = bufio.NewScanner(reader)

func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }
func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = (n * 10) + int(v-'0')
	}
	return
}

type Pair struct {
	a, b int
}

type Pairs []Pair

func (a Pairs) Len() int           { return len(a) }
func (a Pairs) Less(i, j int) bool { return a[i].a < a[j].a }
func (a Pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for t > 0 {
		var n, w, width int
		scanf("%d %d\n", &n, &w)
		count := make([]int, 20)
		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &width)
			} else {
				scanf("%d ", &width)
			}
			count[int(math.Log2(float64(width)))]++
		}
		height, spaceLeft := 1, w
		for i := 0; i < n; i++ {
			largest := -1
			for size := 19; size >= 0; size-- {
				if count[size] > 0 && ((1 << size) <= spaceLeft) {
					largest = size
					break
				}
			}
			if largest == -1 {
				spaceLeft = w
				height++
				for size := 19; size >= 0; size-- {
					if count[size] > 0 && ((1 << size) <= spaceLeft) {
						largest = size
						break
					}
				}
			}

			count[largest] -= 1
			spaceLeft -= (1 << largest)
		}
		printf("%d\n", height)
		t--
	}

}
