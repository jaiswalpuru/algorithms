package main

import (
	"fmt"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func compare_versions(version1, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")

	n, m := len(v1), len(v2)
	p, q := 0, 0
	for p < n && q < m {
		val_1, _ := strconv.Atoi(v1[p])
		val_2, _ := strconv.Atoi(v2[q])
		if val_2 > val_1 {
			return -1
		} else if val_2 < val_1 {
			return 1
		} else {
			p++
			q++
		}
	}

	if q < m {
		for q < m {
			val, _ := strconv.Atoi(v2[q])
			if val != 0 {
				return -1
			}
			q++
		}
	}

	if p < n {
		for p < n {
			val, _ := strconv.Atoi(v1[p])
			if val != 0 {
				return 1
			}
			p++
		}
	}

	return 0
}

func main() {
	fmt.Println(compare_versions("1.0.1", "1"))
}
