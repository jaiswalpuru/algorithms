package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	for t > 0 {
		var s string
		p0, p1 := int(math.MaxInt64), int(math.MaxInt64)
		scanf("%s\n", &s)
		n := len(s)
		posZero, posOne := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			posZero[i] = int(math.MaxInt64)
			posOne[i] = int(math.MaxInt64)
		}
		for i := n - 1; i >= 0; i-- {
			posOne[i] = p1
			posZero[i] = p0
			if s[i] == '1' {
				p1 = i
			} else {
				p0 = i
			}
		}

		if p0 == int(math.MaxInt64) {
			printf("%s\n", "0")
			t--
			continue
		}
		if p1 == int(math.MaxInt64) {
			printf("%s\n", "1")
			t--
			continue
		}
		one, zero, cnt := false, false, 0
		l := make([]int, n)
		for i := n - 1; i >= 0; i-- {
			l[i] = cnt
			if s[i] == '1' {
				one = true
			}
			if s[i] == '0' {
				zero = true
			}
			if one && zero {
				cnt++
				one, zero = false, false
			}
		}
		ans := ""
		curr := p1
		for {
			ans += string(s[curr])
			n1, n0 := posOne[curr], posZero[curr]
			if n0 == int(math.MaxInt64) {
				ans += "0"
				break
			}
			if n1 == int(math.MaxInt64) {
				ans += "1"
				break
			}
			if l[n0] <= l[n1] {
				curr = n0
			} else {
				curr = n1
			}
		}
		printf("%s\n", ans)
		t--
	}
}
