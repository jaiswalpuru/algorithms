package main

import "fmt"

func is_strobogrammtic_number(str string) bool {
	n := len(str)

	i, j := 0, n-1
	for i < j {
		if str[i] == '6' {
			if str[j] != '9' {
				return false
			}
		} else if str[i] == '9' {
			if str[j] != '6' {
				return false
			}
		} else {
			if str[i] != '1' && str[i] != '0' && str[i] != '8' {
				return false
			} else {
				if str[i] != str[j] {
					return false
				}
			}
		}
		i++
		j--
	}
	if i == j {
		if str[i] != '1' && str[i] != '0' && str[i] != '8' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(is_strobogrammtic_number("69"))
}
