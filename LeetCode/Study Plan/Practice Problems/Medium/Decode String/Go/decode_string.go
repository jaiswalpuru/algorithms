package main

import (
	"fmt"
	"strconv"
)

func decode_string(s string) string {
	str := []byte(s)
	n := len(str)
	stack := []byte{}
	for i := 0; i < n; {
		if str[i] == ']' {
			temp := ""
			j := len(stack) - 1
			for stack[j] != '[' {
				temp = string(stack[j]) + temp
				stack = stack[:j]
				j--
			}
			stack = stack[:j]
			j--
			num := ""
			for len(stack) > 0 && (stack[j] >= '0' && stack[j] <= '9') {
				num = string(stack[j]) + num
				stack = stack[:j]
				j--
			}
			val, _ := strconv.Atoi(num)
			t := ""
			for i := 0; i < val; i++ {
				t += temp
			}
			stack = append(stack, []byte(t)...)
		} else {
			stack = append(stack, str[i])
		}
		i++
	}
	return string(stack)
}

func main() {
	fmt.Println(decode_string("3[a]2[bc]"))
}
