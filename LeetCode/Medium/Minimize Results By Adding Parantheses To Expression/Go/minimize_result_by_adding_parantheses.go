package main

import "fmt"
import "strings"
import "strconv"

type Pair struct {
	first,
	second int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimize_result_by_adding_parantheses(str string) string {
	s := strings.Split(str, "+")
	left_pair, right_pair := []Pair{}, []Pair{}

	s1 := s[0]
	s2 := s[1]

	for i := 0; i < len(s1); i++ {
		v1, _ := strconv.Atoi(s1[:i])
		v2, _ := strconv.Atoi(s1[i:])
		left_pair = append(left_pair, Pair{v1, v2})
	}

	for i := 0; i < len(s2); i++ {
		v1, _ := strconv.Atoi(s2[:i])
		v2, _ := strconv.Atoi(s2[i:])
		if v1 == 0 {
			right_pair = append(right_pair, Pair{v2, v1})
		} else {
			right_pair = append(right_pair, Pair{v1, v2})
		}
	}

	v := []int{}
	for i := 0; i < len(left_pair); i++ {
		for j := 0; j < len(right_pair); j++ {
			if len(v) == 0 || (max(1, left_pair[i].first)*(left_pair[i].second+right_pair[j].first)*max(1, right_pair[j].second) < max(1, v[0])*(v[1]+v[2])*max(1, v[3])) {
				v = []int{left_pair[i].first, left_pair[i].second, right_pair[j].first, right_pair[j].second}
			}
		}
	}

	res := ""
	if v[0] != 0 {
		res += strconv.Itoa(v[0])
	}
	res += "(" + strconv.Itoa(v[1]) + "+" + strconv.Itoa(v[2]) + ")"
	if v[3] != 0 {
		res += strconv.Itoa(v[3])
	}
	return res
}

func main() {
	fmt.Println(minimize_result_by_adding_parantheses("12+34"))
	fmt.Println(minimize_result_by_adding_parantheses("247+38"))
}
