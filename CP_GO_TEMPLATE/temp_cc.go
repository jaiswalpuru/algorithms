package main

import (
	"bufio"
	"fmt"
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

func boolIsSubsequence(str1, str2 string, m, n int) bool {
	if m == 0 {
		return true
	}
	if n == 0 {
		return false
	}
	if str1[m-1] == str2[n-1] {
		return boolIsSubsequence(str1, str2, m-1, n-1)
	}
	//if last characters are not matching
	return boolIsSubsequence(str1, str2, m, n-1)

}

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for t > 0 {
		var str string
		scanf("%s\n", &str)
		n := len(str)
		indexOne, indexZero := []int{}, []int{}
		for i := 0; i < n; i++ {
			if str[i] == '1' {
				indexOne = append(indexOne, i)
			} else {
				indexZero = append(indexZero, i)
			}
		}

		if len(indexOne) == 0 {
			printf("%s\n", "1")
		} else if len(indexZero) == 0 {
			printf("%s\n", "0")
		} else {

		}

		t--
	}
}
