/*
Given two version numbers, version1 and version2, compare them.

Version numbers consist of one or more revisions joined by a dot '.'. Each revision consists of digits and may
contain leading zeros. Every revision contains at least one character. Revisions are 0-indexed from left to right,
with the leftmost revision being revision 0, the next revision being revision 1, and so on. For example 2.5.33 and 0.1 are
valid version numbers.

To compare version numbers, compare their revisions in left-to-right order. Revisions are compared using their integer value
ignoring any leading zeros. This means that revisions 1 and 001 are considered equal. If a version number does not specify a
revision at an index, then treat the revision as 0. For example, version 1.0 is less than version 1.1 because their revision
0s are the same, but their revision 1s are 0 and 1 respectively, and 0 < 1.

Return the following:

If version1 < version2, return -1.
If version1 > version2, return 1.
Otherwise, return 0.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compare(ver1, ver2 string) int {
	version_one := strings.Split(ver1, ".")
	version_two := strings.Split(ver2, ".")

	n, m := len(version_one), len(version_two)
	if n > m {
		for i := 0; i < n-m; i++ {
			version_two = append(version_two, "0")
		}
	}
	if m > n {
		for i := 0; i < m-n; i++ {
			version_one = append(version_one, "0")
		}
		n = m
	}

	for i := 0; i < n; i++ {
		if version_one[i] != version_two[i] {
			n1, _ := strconv.Atoi(version_one[i])
			n2, _ := strconv.Atoi(version_two[i])
			fmt.Println(n1, n2)
			if n1 > n2 {
				return 1
			} else if n1 < n2 {
				return -1
			}

		}
	}

	return 0
}

func main() {
	// fmt.Println(compare("1.01", "1.001"))
	// fmt.Println(compare("1.0", "1.0.0"))
	// fmt.Println(compare("0.1", "1.1"))
	fmt.Println(compare("1", "1.1"))
}
