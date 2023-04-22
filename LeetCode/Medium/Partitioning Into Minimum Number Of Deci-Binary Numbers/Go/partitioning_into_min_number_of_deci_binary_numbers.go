package main

import "fmt"

func paritioning_into_min_number_of_binary_numbers(str string) int {
	s := []byte(str)
	max_digit := s[0]
	for i := 0; i < len(s); i++ {
		if max_digit < s[i] {
			max_digit = s[i]
		}
	}

	return int(max_digit - '0')
}

func main() {
	fmt.Println(paritioning_into_min_number_of_binary_numbers("32"))
}
