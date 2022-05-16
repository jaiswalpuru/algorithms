package main

import "fmt"
import "strings"
import "strconv"

//refer LC soln

func restore_ip_address(s string) []string {
	n := len(s)
	res := []string{}
	backtrack(s, n, -1, 3, []string{}, &res)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func update_res(s string, n int, cur_pos int, set []string, res *[]string) {
	seg := s[cur_pos+1 : n]
	if is_valid(seg) {
		set = append(set, seg)
		*res = append(*res, strings.Join(set, "."))
		set = set[:len(set)-1]
	}
}

func backtrack(s string, n int, prev_pos int, dots int, set []string, res *[]string) {
	max_pos := min(n-1, prev_pos+4)

	for cur_pos := prev_pos + 1; cur_pos < max_pos; cur_pos++ {
		seg := s[prev_pos+1 : cur_pos+1]
		if is_valid(seg) {
			//place the dot
			set = append(set, seg)
			if dots-1 == 0 {
				//if all 3 dots have been placed add it to the solution
				update_res(s, n, cur_pos, set, &(*res))
			} else {
				backtrack(s, n, cur_pos, dots-1, set, res)
			}
			//remove the dot
			set = set[:len(set)-1]
		}
	}
}

func is_valid(s string) bool {
	n := len(s)
	if n > 3 {
		return false
	}
	if s[0] != '0' {
		val, _ := strconv.Atoi(s)
		if val <= 255 {
			return true
		}
		return false
	} else {
		return n == 1
	}
}

func main() {
	fmt.Println(restore_ip_address("25525511135"))
}
