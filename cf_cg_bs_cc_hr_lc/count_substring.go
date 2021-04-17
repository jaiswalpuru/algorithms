package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
	far    = make([]int, 100000+1)
	sumfar = make([]int, 100000+1)
)

func main() {
	defer writer.Flush()

	var (
		T, N, K, Q int    // n -> len of string , q -> queries, o/p: count the number of strings with atmost k characters
		str        string //len n
		L, R       int    //for each query
		cnt1, cnt0 = 0, 0
	)
	scanf("%d\n", &T)

	for ; T > 0; T-- {

		scanf("%d %d %d\n", &N, &K, &Q)
		scanf("%s\n", &str) // input string
		str = " " + str
		stRune := []rune(str)

		j := 1
		if stRune[1] == '1' {
			cnt1++
		} else {
			cnt0++
		}

		for i := 1; i <= N; i++ {
			for j <= N && cnt0 <= K && cnt1 <= K {
				j++
				if j > N {
					break
				}
				if stRune[j] == '0' {
					cnt0++
				} else {
					cnt1++
				}
			}
			far[i] = j

			if stRune[i] == '1' {
				cnt1--
			} else {
				cnt0--
			}
		}

		sumfar[0] = 0
		for i := 1; i <= N; i++ {
			sumfar[i] = sumfar[i-1] + far[i]
		}

		ans := 0
		for ; Q > 0; Q-- {
			scanf("%d %d\n", &L, &R)
			writer.Flush()
			k1 := L - 1
			k2 := R + 1
			var km int
			for k2-k1 > 1 {
				km = (k2 + k1) / 2
				if far[km] <= R {
					k1 = km
				} else {
					k2 = km
				}
			}
			km = k1
			ans = sumfar[km] - sumfar[L-1]
			ans += ((R - km) * (R + 1)) - ((R*(R+1))/2 - (L*(L-1))/2)
			printf("%d\n", ans)
			writer.Flush()
		}
	}
}
