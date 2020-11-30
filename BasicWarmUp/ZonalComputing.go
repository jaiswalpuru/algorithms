package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {

	var (
		n, temp, val, ans int
		arr               = []int{}
		sc                = bufio.NewScanner(os.Stdin)
	)

	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	temp = n

	for n > 0 {
		sc.Scan()
		val, _ = strconv.Atoi(sc.Text())
		arr = append(arr, val)
		n--
	}

	sort.Ints(arr)
	ans = 0
	n = temp

	for i := 0; i < n; i++ {
		temp = arr[i] * (n - i)
		ans = int(math.Max(float64(ans), float64(temp)))
	}

	fmt.Println(ans)

}
