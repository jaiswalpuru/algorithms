//Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
//For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.

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

func countDecodings(r []rune, size int) int {
	if size == 0 || size == 1 {
		return 1
	}
	if r[0] == '0' {
		return 0
	}
	count := 0

	if r[size-1] != '0' {
		count = countDecodings(r, size-1)
	}

	if r[size-2] == '1' || (r[size-2] == '2' && r[size-1] < '7') {
		count += countDecodings(r, size-2)
	}
	return count
}

func main() {
	defer writer.Flush()

	var s string
	scanf("%s\n", &s)
	sRune := []rune(s)

	printf("%d\n", countDecodings(sRune, len(sRune)))
}
