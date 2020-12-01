package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	t = toInt([]byte(sc.Text()))

	for t > 0 {

		var (
			str       string
			d0, d1, k int
		)

		sc.Scan()
		str = sc.Text()
		vals := strings.Split(str, " ")
		k = toInt([]byte(vals[0]))
		d0 = toInt([]byte(vals[1]))
		d1 = toInt([]byte(vals[2]))

		res := d0 + d1 //d2
		cycle := ((2 * res) % 10) + ((4 * res) % 10) + ((8 * res) % 10) + ((6 * res) % 10)
		sum := d0 + d1 + (d0+d1)%10 + ((k-3)/4)*cycle
		extra := (k - 3) % 4
		temp := 0

		if extra == 1 {
			temp = (2 * res) % 10
		} else if extra == 2 {
			temp = ((2 * res) % 10) + ((4 * res) % 10)
		} else if extra == 3 {
			temp = ((2 * res) % 10) + ((4 * res) % 10) + ((8 * res) % 10)
		}

		sum += temp

		if sum%3 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--

	}

}
