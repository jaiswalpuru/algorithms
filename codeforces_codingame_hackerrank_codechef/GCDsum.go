package main

import (
	"bufio"
	"fmt"
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

func gcd(a, b int64) int64 {
	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func gcdSum(n int64) int64 {
	temp, digit := n, int64(0)
	for temp > 0 {
		digit += temp % 10
		temp = temp / 10
	}
	ans := gcd(n, digit)
	return ans
}

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for t > 0 {
		var n int64
		scanf("%d\n", &n)
		if gcdSum(n) != 1 {
			printf("%d\n", n)
		} else if gcdSum(n+1) != 1 {
			printf("%d\n", n+1)
		} else if gcdSum(n+2) != 1 {
			printf("%d\n", n+2)
		}
		t--
	}

}
