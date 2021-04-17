package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }

func main() {

	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {

		var str string
		scanf("%s\n", &str)

		arr := []rune(str)
		lt, rt := 0, 0

		for i := 0; i < len(str); i++ {
			if arr[i] == '<' {
				lt++
			} else {
				lt--
			}
			if lt < 0 {
				break
			}
			if lt == 0 {
				rt = i + 1
			}
		}

		fmt.Println(rt)

	}

}
