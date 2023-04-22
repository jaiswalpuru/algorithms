/*
Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.
*/

package main

import "fmt"

var mp = map[int]string{
	1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H", 9: "I", 10: "J", 11: "K", 12: "L", 13: "M", 14: "N", 15: "O",
	16: "P", 17: "Q", 18: "R", 19: "S", 20: "T", 21: "U", 22: "V", 23: "W", 24: "X", 25: "Y", 26: "Z",
}

func get_column(n int) string {
	str := ""
	for n > 0 {
		if n%26 == 0 {
			str = mp[26] + str
			n = n/26 - 1
		} else {
			str = mp[n%26] + str
			n = n / 26
		}
	}
	return str
}

func main() {
	fmt.Println(get_column(701))
	fmt.Println(get_column(52))
	fmt.Println(get_column(10000))
}
