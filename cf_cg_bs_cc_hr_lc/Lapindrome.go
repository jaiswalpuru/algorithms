package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		t   int
		str string
		sc  = bufio.NewScanner(os.Stdin)
	)

	sc.Scan()
	t, _ = strconv.Atoi(sc.Text())

	for t > 0 {

		var (
			f       = 0
			strRune []rune
			inRu    = make(map[rune]int)
		)

		sc.Scan()
		str = sc.Text()
		strRune = []rune(str)

		if len(str) == 1 {
			fmt.Println("NO")
			return
		}

		for i := 0; i < len(strRune)/2; i++ {
			inRu[strRune[i]]++
		}

		if len(str)%2 == 1 {

			for i := len(strRune)/2 + 1; i < len(strRune); i++ {
				inRu[strRune[i]]--
			}

		} else {

			for i := len(strRune) / 2; i < len(strRune); i++ {
				inRu[strRune[i]]--
			}

		}

		for _, val := range inRu {
			if val != 0 {
				f = 1
				break
			}
		}

		if f == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}

}
