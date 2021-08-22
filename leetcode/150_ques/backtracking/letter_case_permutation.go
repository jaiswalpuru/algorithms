/*
Given a string s, we can transform every letter individually to be lowercase or uppercase to create another string.

Return a list of all possible strings we could create. You can return the output in any order.

Example 1:
Input: s = "a1b2"
Output: ["a1b2","a1B2","A1b2","A1B2"]

Example 2:
Input: s = "3z4"
Output: ["3z4","3Z4"]

Example 3:
Input: s = "12345"
Output: ["12345"]

Example 4:
Input: s = "0"
Output: ["0"]
*/

package main

import "fmt"

var (
	res = []string{}
)

func _all_perm(str []byte, s string, res *[]string, i int) {
	if len(str) == len(s) {
		*res = append(*res, string(str))
		return
	}
	if i < len(s) {
		_all_perm(append(str, s[i]), s, res, i+1)
		if s[i] >= 'a' && s[i] <= 'z' {
			_all_perm(append(str, s[i]-32), s, res, i+1)
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			_all_perm(append(str, s[i]+32), s, res, i+1)
		}
	}
}

func all_perm(str string) []string {
	res := []string{}
	_all_perm([]byte{}, str, &res, 0, "")
	return res
}

func main() {
	fmt.Println(all_perm("3z4"))
	fmt.Println(all_perm("a1b2"))
	fmt.Println(all_perm("12345"))
	fmt.Println(all_perm("0"))
}
