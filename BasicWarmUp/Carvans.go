package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//Used for fast scan
func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {

	var (
		t  int
		sc = bufio.NewScanner(os.Stdin)
	)

	sc.Scan()
	t = toInt(sc.Bytes())

	for t > 0 {

		var n int
		sc.Scan()
		n = toInt(sc.Bytes())
		arr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		cnt := 0
		temp := math.MaxInt32

		for i := 0; i < n; i++ {
			if arr[i] <= temp {
				cnt++
				temp = arr[i]
			}
		}

		fmt.Println(cnt)
		t--
	}
}
