package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var scanner *bufio.Scanner = bufio.NewScanner(reader)

const mod = 1000000007

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

//define constants here
const (
	N = 1001
	K = 1001
)

//define global variables here
var (
	dp   = [N][K][2]int{}
	n, k int
)

//Write your code from here

func solve(cur, k, dir int) int {

	if k == 1 {
		return 1
	}

	if dp[cur][k][dir] != -1 {
		return dp[cur][k][dir]
	}

	ans := 2 //me and my copy
	//dir are two make copy forward and one backward
	if dir == 1 {

		if cur < n {
			ans += solve(cur+1, k, dir) - 1
		}
		ans %= mod
		if cur > 1 {
			ans += solve(cur-1, k-1, 1-dir) - 1
		}

		ans %= mod
		dp[cur][k][dir] = ans

	} else {

		if cur > 1 {
			ans += solve(cur-1, k, dir) - 1
		}
		ans %= mod
		if cur < n {
			ans += solve(cur+1, k-1, 1-dir) - 1
		}
		ans %= mod
		dp[cur][k][dir] = ans
	}
	return ans

}

func main() {
	defer writer.Flush()
	var t int
	scanf("%d\n", &t)

	for t > 0 {
		scanf("%d %d\n", &n, &k)

		for i := 0; i < 1001; i++ {
			for j := 0; j < 1001; j++ {
				dp[i][j][0] = -1
				dp[i][j][1] = -1
			}
		}
		printf("%d\n", solve(1, k, 1))
		t--
	}

}
