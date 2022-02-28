package main

import (
	"fmt"
	"strconv"
)

func summary_range(arr []int) []string {
	res := []string{}
	n := len(arr)
	if n == 0 {
		return nil
	}
	l, h := arr[0], arr[0]
	str := ""
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == 1 {
			h = arr[i]
		} else {
			if l != h {
				str += strconv.Itoa(l) + "->" + strconv.Itoa(h)
				res = append(res, str)
				str = ""
			} else {
				res = append(res, strconv.Itoa(l))
			}
			l = arr[i]
			h = arr[i]
		}
	}
	if l == h {
		str = strconv.Itoa(l)
	} else {
		str = strconv.Itoa(l) + "->" + strconv.Itoa(h)
	}

	if len(res) == 0 || res[len(res)-1] != str {
		res = append(res, str)
	}
	return res
}

func main() {
	fmt.Println(summary_range([]int{0, 1, 2, 4, 5, 7}))
}
