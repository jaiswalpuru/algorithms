package main

import "fmt"

func addToArrayForm(num []int, k int) []int {
	size := len(num)
	res := []int{}
	i, carry := size-1, 0

	for i >= 0 && k > 0 {
		sum := num[i] + k%10 + carry
		res = append(res, sum%10)
		carry = sum / 10
		k /= 10
		i--
	}

	for i >= 0 {
		sum := num[i] + carry
		res = append(res, sum%10)
		carry = sum / 10
		i--
	}

	for k > 0 {
		sum := (k % 10) + carry
		res = append(res, sum%10)
		k = k / 10
		carry = sum / 10
	}

	for carry > 0 {
		res = append(res, carry%10)
		carry /= 10
	}

	return reverse(res)
}

func reverse(a []int) []int {
	size := len(a)
	for i := 0; i < size/2; i++ {
		a[i], a[size-1-i] = a[size-1-i], a[i]
	}
	return a
}

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
}
