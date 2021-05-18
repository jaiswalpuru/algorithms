/*
Given a string of length n and and integer k. The string can be manipulated by taking one of the first k letters
and moving it to the end.
Wap to determine the lexicographically smallest string that can be created after an unlimited number of moves.
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	str := "daily"
	k := 2

	if k == 1 {
		temp := str
		for i := 1; i < len(str); i++ {
			if str[i:]+str[:i] < temp {
				temp = str[i:] + str[:i]
			}
		}
		fmt.Println(temp)
	} else {
		temp := strings.Split(str, "")
		sort.Strings(temp)
		fmt.Println(strings.Join(temp, ""))
	}

}
