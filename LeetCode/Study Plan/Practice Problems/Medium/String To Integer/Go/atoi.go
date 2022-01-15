package main

import (
	"fmt"
	"math"
)

func atoi(str string) int {
	sign := 1
	ans := 0
	ind := 0
	n := len(str)

	//discarding the spaces in the start
	for ind < n && str[ind] == ' ' {
		ind++
	}

	//check for signs
	if ind < n && str[ind] == '+' {
		sign = 1
		ind++
	} else if ind < n && str[ind] == '-' {
		sign = -1
		ind++
	}

	for ind < n && (str[ind] >= '0' && str[ind] <= '9') {
		num := int(str[ind] - '0')

		if (ans > math.MaxInt32/10) || (ans == math.MaxInt32/10 && num > math.MaxInt32%10) {
			return func() int {
				if sign == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
			}()
		}

		ans = 10*ans + num
		ind++
	}
	return sign * ans
}

func main() {
	// fmt.Println(atoi("   -42"))
	fmt.Println(atoi("91283472332"))
	fmt.Println(atoi("-9223372036854775808"))
}
