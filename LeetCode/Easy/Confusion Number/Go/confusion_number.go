package main

import "fmt"

var mp = map[int]int{
	0: 0, 1: 1, 6: 9, 8: 8, 9: 6,
}

func confusion_number(n int) bool {
	res := 0
	num := n
	for n > 0 {
		v := n % 10
		if _, ok := mp[v]; !ok {
			return false
		}
		res = (res * 10) + mp[v]
		n /= 10
	}
	return res != num
}

func main() {
	fmt.Println(confusion_number(6))
}
