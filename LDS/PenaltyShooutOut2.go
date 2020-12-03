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

		var (
			a, b, teamA, teamB = 0, 0, 0, 0
			n                  int
			str                string
		)

		scanf("%d\n", &n)
		scanf("%s\n", &str)
		arr := []rune(str)
		limit := 2 * n

		for i := 0; i < limit; i++ {

			if i%2 == 0 {
				//teamA
				teamA += int(arr[i]) - 48
				a++
			} else {
				//teamB
				teamB += int(arr[i]) - 48
				b++
			}

			if (teamA-teamB > n-b) || (teamB-teamA > n-a) {
				printf("%d\n", i+1)
				break
			}

		}

		if teamA == teamB {
			printf("%d\n", n*2)
		}

	}

}
