/*
Your friend is typing his name into a keyboard. Sometimes, when typing a character c, the key might get long pressed,
and the character will be typed 1 or more times.

You examine the typed characters of the keyboard. Return True if it is possible that it was your friends name, with some
characters (possibly none) being long pressed.
*/
package main

import "fmt"

func long_pressed(name, typed string) bool {
	n, m := len(name), len(typed)
	if m < n {
		return false
	}

	if n == m {
		if name != typed {
			return false
		}
	}

	i, j := 0, 0
	for i = 0; i < n; {
		if j < m && name[i] == typed[j] {
			cnt_1, cnt_2 := 0, 0
			for i < n-1 && name[i] == name[i+1] {
				cnt_1++
				i++
			}

			for j < m-1 && typed[j] == typed[j+1] {
				cnt_2++
				j++
			}

			i++
			j++
			if cnt_1 > cnt_2 {
				return false
			}
		} else {
			return false
		}
	}

	if i < n || j < m {
		return false
	}
	return true
}

func main() {
	//fmt.Println(long_pressed("alex", "aaleexa"))
	fmt.Println(long_pressed("pyplrz", "ppyypllr"))
}
