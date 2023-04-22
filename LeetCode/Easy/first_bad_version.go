/*
You are a product manager and currently leading a team to develop a new product. Unfortunately,
the latest version of your product fails the quality check. Since each version is developed based
on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes
all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a
function to find the first bad version. You should minimize the number of calls to the API.
*/

package main

import "fmt"

//junk function does nothing
func isBadVersion(n int) bool {
	return true
}

func first_bad_version(n int) int {
	l, h := 1, n
	for l < h {
		mid := (l + h) / 2
		if isBadVersion(mid) {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(first_bad_version(5))
}
