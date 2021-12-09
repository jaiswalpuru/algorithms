/*
Your friend is typing his name into a keyboard. Sometimes, when typing a character c,
the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard. Return True if it is possible that it was
your friends name, with some characters (possibly none) being long pressed.

Example 1:
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.

Example 2:
Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.

Example 3:
Input: name = "leelee", typed = "lleeelee"
Output: true

Example 4:
Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.
*/

package main

import "fmt"

func longed_pressed_name(name, typed string) bool {
	n, m := len(name), len(typed)
	if m < n {
		return false
	}
	i, j := 0, 0
	if name[0] != typed[0] {
		return false
	}
	i, j = i+1, j+1

	for i < n && j < m {
		if name[i] == typed[j] {
			i++
			j++
		} else {
			if typed[j] == typed[j-1] {
				j++
			} else {
				return false
			}
		}
	}

	if i != n {
		return false
	}
	if j != m {
		for k := j; k < m; k++ {
			if typed[k] != typed[k-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	//fmt.Println(longed_pressed_name("alex", "aaleexxa"))
	//fmt.Println(longed_pressed_name("saeed", "ssaaedd"))
	//fmt.Println(longed_pressed_name("leelee", "lleeelee"))
	//fmt.Println(longed_pressed_name("alex", "aaleex"))
	fmt.Println(longed_pressed_name("pyplrz", "ppyypllr"))
}
