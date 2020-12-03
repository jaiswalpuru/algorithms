package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//arr[0] = Si stores that offer food of ith type
//arr[1] = Vi price of one type of food of that type
//arr[2] = Pi people come to buy that type of food

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
	t, _ = strconv.Atoi(sc.Text())

	for t > 0 {

		var (
			n, ans int
		)

		sc.Scan()
		n, _ = strconv.Atoi(sc.Text())
		arr := make([][3]int, n)
		ans = 0

		for i := 0; i < n; i++ {
			sc.Scan()
			str := sc.Text()
			temp := strings.Split(str, " ")
			arr[i][0] = toInt([]byte(temp[0]))
			arr[i][1] = toInt([]byte(temp[1]))
			arr[i][2] = toInt([]byte(temp[2]))
			arr[i][0]++
		}

		for i := 0; i < n; i++ {
			ans = int(math.Max(float64(ans), float64(arr[i][2]*(arr[i][1]/arr[i][0]))))
		}

		fmt.Println(ans)

		t--
	}
}
