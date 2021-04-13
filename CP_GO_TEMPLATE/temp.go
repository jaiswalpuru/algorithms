package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReaderSize(os.Stdin, 1024*1024)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }
func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }

type Pair struct {
	x, y int
}

type Pairs []Pair

func (a Pairs) Len() int           { return len(a) }
func (a Pairs) Less(i, j int) bool { return a[i].x < a[j].x }
func (a Pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func calcPrefixSum(arr, psa [][]int64, n, m int64) [][]int64 {
	var i, j int64 = 0, 0
	psa[0][0] = arr[0][0]

	//first row
	for i = 1; i < m; i++ {
		psa[0][i] = psa[0][i-1] + arr[0][i]
	}
	//first column
	for i = 1; i < n; i++ {
		psa[i][0] = psa[i-1][0] + arr[i][0]
	}

	//updating rest of the values
	for i = 1; i < n; i++ {
		for j = 1; j < m; j++ {
			psa[i][j] = psa[i-1][j] + psa[i][j-1] - psa[i-1][j-1] + arr[i][j]
		}
	}
	return psa
}

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for t > 0 {
		var k, n, m, ans int64
		scanf("%d %d %d\n", &n, &m, &k)
		arr := make([][]int64, n)
		psa := make([][]int64, n)
		ans = 0
		for i := int64(0); i < n; i++ {
			arr[i] = make([]int64, m)
			psa[i] = make([]int64, m)
			for j := int64(0); j < m; j++ {
				if j == m-1 {
					scanf("%d\n", &arr[i][j])
				} else {
					scanf("%d ", &arr[i][j])
				}
				if arr[i][j] >= k {
					ans++
				}
			}
		}

		l, maxDimension := int64(2), int64(math.Min(float64(m), float64(n)))

		for l < maxDimension {

			for i := int64(0); i < n; i++ {
				for j := int64(0); j < m; j++ {

					temp := int64(0)

					for k := i; k < i+l; k++ {
						for q := j; q < j+l; q++ {
							printf("here break q=%d\n", q)
							temp += arr[k][q]
							printf("%d k=%d q=%d i=%d j=%d\n", arr[k][q], k, q, i, j)
						}
					}

					printf("%s\n", "")
					if float64(temp)/float64(k) >= float64(k) {
						ans++
					}
				}
			}
			l++
		}
		printf("%d\n", ans)
		t--
	}
}
