package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }
func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }

type i64 int64

func main() {

	var t int
	scanf("%d\n", &t)

	for t > 0 {

		//get the inputs
		var n, m, i i64
		scanf("%d %d\n", &n, &m)
		disks := make([]i64, n)
		ques := make([]i64, m)
		for i = 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &disks[i])
			} else {
				scanf("%d ", &disks[i])
			}
		}
		for i = 0; i < m; i++ {
			if i == m-1 {
				scanf("%d\n", &ques[i])
			} else {
				scanf("%d ", &ques[i])
			}
		}

		//solve

		t--
	}
}
