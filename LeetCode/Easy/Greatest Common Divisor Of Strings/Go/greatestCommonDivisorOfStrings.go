package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
}
