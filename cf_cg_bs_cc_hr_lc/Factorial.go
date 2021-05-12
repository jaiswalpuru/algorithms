package main

import (
	"bufio"
	"fmt"
	"os"
)

func toInt(buf []byte) (n int) {

	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return

}

func main() {

	var (
		t, n int
		sc   = bufio.NewScanner(os.Stdin)
	)

	sc.Scan()
	t = toInt(sc.Bytes())

	for ; t > 0; t-- {

		sc.Scan()
		n = toInt(sc.Bytes())
		ans := 0

		for i := 1; n != 0; i++ {
			ans += n / 5
			n = n / 5
		}

		fmt.Println(ans)

	}

}
