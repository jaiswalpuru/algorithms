package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	i := 0
	j := 0
	size := len(chars)

	for i < size {
		cnt := 0
		chars[j] = chars[i]
		for i < size && chars[i] == chars[j] {
			cnt++
			i++
		}

		if cnt > 1 {
			str := []byte(strconv.Itoa(cnt))
			for k := 0; k < len(str); k++ {
				j++
				chars[j] = str[k]
			}
		}
		j++
	}

	return j
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}
