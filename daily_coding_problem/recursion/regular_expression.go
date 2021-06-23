/*
Implement regular expression matching with the following special characters:

-> . (period) which matches any single character
-> * (asterisk) which matches zero or more of the preceding element
That is, implement a function that takes in a string and a valid regular expression and returns whether
or not the string matches the regular expression.

For example, given the regular expression "ra." and the string "ray", your function should return true.
The same regular expression on the string "raymond" should return false.

Given the regular expression ".*at" and the string "chat", your function should return true.
The same regular expression on the string "chats" should return false.
*/

package main

import "fmt"

func match_first_character(s, r string) bool { return (r[0] == s[0]) || (len(s) > 0 && r[0] == '.') }

//Note : '*' matches zero or more of the preceeding element
func match(s, r string) bool {

	//check if the string is empty
	if r == "" {
		return s == ""
	}

	//check if the length of regex is 1 or the first characted is not succeeded by '*'
	if len(r) == 1 || r[1] != []byte("*")[0] {
		if match_first_character(s, r) {
			return match(s[1:], r[1:])
		} else {
			return false
		}
	} else {
		//the first character is succeeded by a '*'
		if match(s, r[2:]) {
			return true
		}

		i := 0
		for match_first_character(s[i:], r) {
			if match(s[i+1:], r[2:]) {
				return true
			}
			i += 1
		}
	}
	return false
}

func main() {
	regExp := ".*at"
	str := "chat"
	fmt.Println(str, regExp, match(str, regExp))
	regExp = "ra."
	str = "ray"
	fmt.Println(str, regExp, match(str, regExp))
}
