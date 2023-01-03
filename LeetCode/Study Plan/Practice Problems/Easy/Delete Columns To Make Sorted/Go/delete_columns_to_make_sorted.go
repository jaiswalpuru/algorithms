package main

import (
	"fmt"
	"strings"
)

func delete_columns_to_make_sorted(strs []string) int {
	str := strings.Join(strs, "")
	cnt := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := i + len(strs[0]); j < len(str); j += len(strs[0]) {
			if str[j-len(strs[0])] > str[j] {
				cnt++
				break
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(delete_columns_to_make_sorted([]string{"zyx", "wvu", "tsr"}))
}
