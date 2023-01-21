package main

import (
	"fmt"
	"strconv"
	"strings"
)

//refer LC soln
func restoreIpAddresses(s string) []string {
	res := []string{}
	backtrack(s, -1, 3, &res, []string{})
	return res
}

func backtrack(s string, prev, dots int, res *[]string, temp []string) {
	for curr := prev + 1; curr < min(len(s)-1, prev+4); curr++ {
		seg := s[prev+1 : curr+1]
		if isValid(seg) {
			temp = append(temp, seg)
			if dots-1 == 0 {
				updateRes(s, curr, temp, res)
			} else {
				backtrack(s, curr, dots-1, res, temp)
			}
			temp = temp[:len(temp)-1]
		}
	}
}

func updateRes(s string, curr int, temp []string, res *[]string) {
	seg := s[curr+1 : len(s)]
	if isValid(seg) {
		temp = append(temp, seg)
		*res = append(*res, strings.Join(temp, "."))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func isValid(s string) bool {
	if len(s) > 3 {
		return false
	}
	if s[0] != '0' {
		val, _ := strconv.Atoi(s)
		if val <= 255 {
			return true
		}
		return false
	}
	return len(s) == 1
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
