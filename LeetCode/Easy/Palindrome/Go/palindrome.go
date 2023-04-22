package main

import (
	"fmt"
)

func palindrome(n int) bool {
	if (n%10 == 0 && n != 0) || n < 0 {
		return false
	}

	res := 0
	i := 10
	temp := n
	for n > 0 {
		res = (res * i) + (n % 10)
		fmt.Println(res)
		n = n / 10
	}
	return temp == res
}

func main() {
	fmt.Println(palindrome(10))
	// fmt.Println(palindrome(-121))

}
