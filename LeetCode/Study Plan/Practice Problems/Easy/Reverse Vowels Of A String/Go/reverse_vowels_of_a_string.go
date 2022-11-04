package main

import "fmt"

func reverse_vowels_of_a_string(s string) string {
	vow := []byte{}
	str := []byte(s)
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'o' || str[i] == 'u' || str[i] == 'i' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {
			vow = append(vow, str[i])
		}
	}
	reverse(&vow)
	k := 0
	for i := 0; i < n; i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'o' || str[i] == 'u' || str[i] == 'i' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {
			str[i] = vow[k]
			k++
		}
	}
	return string(str)
}

func reverse(vow *[]byte) {
	n := len(*vow)
	for i := 0; i < n/2; i++ {
		(*vow)[i], (*vow)[n-1-i] = (*vow)[n-1-i], (*vow)[i]
	}
}

func main() {
	fmt.Println(reverse_vowels_of_a_string("hello"))
}
