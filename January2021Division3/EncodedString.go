package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	ascii  = []rune{97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112}
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

func getRune(r, temp []rune, val int) rune {
	if len(temp) == 1 {
		return temp[0]
	}
	if r[0] == '0' {
		return getRune(r[1:], temp[:val], val/2)
	}
	return getRune(r[1:], temp[val:], val/2)
}

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {
		var (
			n   int
			str string
			ans = []rune{}
		)
		scanf("%d\n%s\n", &n, &str)
		r := []rune(str)

		//The first bit (from the left) of the code is 0 if the letter lies among the first 8 letters
		//else it is 1, signifying that it lies among the last 8 letters
		//The second bit of the code is 0 if the letter lies among the first 4 letters of those 8 letters found in the previous step
		//else it's 1, signifying that it lies among the last 4 letters of those 8 letters
		for i := 0; i < n; i += 4 {
			ans = append(ans, getRune(r[i:i+4], ascii, 8))
		}
		fmt.Println(string(ans))
	}
}
