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

func main() {
	defer writer.Flush()

	var s string
	scanf("%s\n", &s)
	n, x1, x2, y1, y2 := len(s), -1, -1, -1, -1
	for i := 0; i < n && (i+1) < n; i++ {
		if s[i:i+2] == "AB" {
			x1 = i
			break
		}
	}
	for j := x1 + 2; j < n && (j+1) < n; j++ {
		if s[j:j+2] == "BA" {
			x2 = j
			break
		}
	}
	for i := 0; i < n && (i+1) < n; i++ {
		if s[i:i+2] == "BA" {
			y1 = i
			break
		}
	}
	for j := y1 + 2; j < n && (j+1) < n; j++ {
		if s[j:j+2] == "AB" {
			y2 = j
			break
		}
	}

	if (x1 != -1 && x2 != -1) || (y1 != -1 && y2 != -1) {
		printf("%s\n", "YES")
	} else {
		printf("%s\n", "NO")
	}
}
