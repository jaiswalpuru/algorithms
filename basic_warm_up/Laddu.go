package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//var reader *bufio.Reader = bufio.NewReader(os.Stdin)
//var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

//func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
//func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

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

	fmt.Scanf("%d\n", &t)

	for t > 0 {

		var (
			str, origin string
			bonus, ans  = 0, 0
			activity    int
		)

		sc.Scan()
		vals := strings.Split(sc.Text(), " ")
		activity, _ = strconv.Atoi(vals[0])
		origin = vals[1]

		for i := 0; i < activity; i++ {

			sc.Scan()
			str = sc.Text()
			arr := strings.Split(str, " ")

			switch temp := arr[0]; temp {
			case "CONTEST_WON":
				t1 := toInt([]byte(arr[1]))
				if t1 < 20 {
					bonus = 20 - t1
				}
				ans += 300 + bonus
			case "TOP_CONTRIBUTOR":
				ans += 300
			case "BUG_FOUND":
				t1 := toInt([]byte(arr[1]))
				ans += t1
			case "CONTEST_HOSTED":
				ans += 50
			}

			bonus = 0

		}

		if strings.Compare(origin, "INDIAN") == 0 {
			fmt.Printf("%d\n", ans/200)
		} else {
			fmt.Printf("%d\n", ans/400)
		}

		t--

	}
}
