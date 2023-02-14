package main

import "fmt"

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	res := []byte{}

	for i >= 0 && j >= 0 {
		if a[i] == '1' && b[j] == '1' {
			if carry != 0 {
				res = append(res, '1')
			} else {
				res = append(res, '0')
			}
			carry = 1
		} else {
			if carry != 0 {
				if a[i] == '0' && b[j] == '0' {
					res = append(res, '1')
					carry = 0
				} else {
					res = append(res, '0')
					carry = 1
				}
			} else {
				if a[i] == '0' && b[j] == '0' {
					res = append(res, '0')
				} else {
					res = append(res, '1')
				}
			}
		}
		i--
		j--
	}

	for i >= 0 {
		if carry == 0 {
			res = append(res, a[i])
		} else {
			if a[i] == '1' {
				res = append(res, '0')
				carry = 1
			} else {
				res = append(res, '1')
				carry = 0
			}
		}
		i--
	}

	for j >= 0 {
		if carry == 0 {
			res = append(res, b[j])
		} else {
			if b[j] == '1' {
				res = append(res, '0')
				carry = 1
			} else {
				res = append(res, '1')
				carry = 0
			}
		}
		j--
	}

	if carry != 0 {
		res = append(res, '1')
	}

	return string(reverse(res))
}

func reverse(a []byte) []byte {
	size := len(a)
	for i := 0; i < size/2; i++ {
		a[i], a[size-1-i] = a[size-1-i], a[i]
	}
	return a
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
