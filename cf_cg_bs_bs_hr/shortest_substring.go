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

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)
	for t > 0 {
		var s string
		scanf("%s\n", &s)

		ans := ""
		if len(s) <= 3 {
			printf("%s\n", s)
		} else {
			ans += string(s[0])
			for i := 1; i < len(s)-1; i += 2 {
				ans += string(s[i])
			}
			ans += string(s[len(s)-1])
			printf("%s\n", ans)
		}
		t--
	}
}
