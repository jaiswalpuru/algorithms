/*
	Given a string s and an integer k, break up the string into multiple lines such that each line has a
	length of k or less. You must break it up so that words don't break across lines.
	Each line has to have the maximum possible amount of words. If there's no way to break the text up,
	then return null.

	You can assume that there are no spaces at the ends of the string and that there is exactly
	one space between each word.

	For example, given the string "the quick brown fox jumps over the lazy dog" and k = 10,
	you should return: ["the quick", "brown fox", "jumps over", "the lazy", "dog"].
	No string in the list has a length of more than 10.
*/

package main

import (
	"fmt"
	"strings"
)

func segregate(str string, k int) {
	res := []string{}
	str_arr := strings.Split(str, " ")
	temp := str_arr[0]
	for i := 1; i < len(str_arr); i++ {
		if len(temp)+len(str_arr[i])+1 > k {
			res = append(res, temp)
			temp = str_arr[i]
		} else {
			if temp != "" {
				temp += " " + str_arr[i]
			} else {
				temp += str_arr[i]
			}
		}
	}
	if temp != "" {
		res = append(res, temp)
	}
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func main() {

	str := "the quick brown fox jumps over the lazy dog"
	k := 10

	segregate(str, k)
}
