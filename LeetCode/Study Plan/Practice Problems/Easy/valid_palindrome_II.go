/*
Given a string s, return true if the s can be palindrome after deleting at most one character from it.
*/

package main

import "fmt"

func reverse_string(str string) string {
	arr := []byte(str)
	n := len(str)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return string(arr)
}

func valid_palindrome(s string) bool {

	n := len(s)
	if n == 1 {
		return true
	}

	i, j := 0, n-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			temp := s[0:i] + s[i+1:]
			rev := reverse_string(temp)
			if temp == rev {
				return true
			}
			temp = s[0:j] + s[j+1:]
			rev = reverse_string(temp)
			if temp == rev {
				return true
			}
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(valid_palindrome("ebcbbececabbacecbbcbe"))
	fmt.Println(valid_palindrome("abca"))
	fmt.Println(valid_palindrome("abc"))
}
