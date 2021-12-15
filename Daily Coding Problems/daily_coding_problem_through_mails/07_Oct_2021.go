/*
Determine whether there exists a one-to-one character mapping from one string s1 to another s2.

For example, given s1 = abc and s2 = bcd, return true since we can map a to b, b to c, and c to d.

Given s1 = foo and s2 = bar, return false since the o cannot map to two characters.
*/

package main

import "fmt"

func one_to_one_mapping(str_one, str_two string) bool {
	mp := make(map[byte]byte)

	i, j := 0, 0 //starting index for the str_one and str_two
	n, m := len(str_one), len(str_two)

	for i < n && j < m {
		if _, ok := mp[str_one[i]]; ok {
			return false
		}
		mp[str_one[i]] = str_two[j]
		i++
		j++
	}
	return true
}

func main() {
	fmt.Println(one_to_one_mapping("abc", "bcd"))
	fmt.Println(one_to_one_mapping("foo", "bar"))
}
