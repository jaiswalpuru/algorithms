package main

import "fmt"

type Pair struct {
	c byte
	cnt int
}

func remove_all_duplicates_in_string(str string, k int) string {
	s := []Pair{}
	n := len(str)

	for i:=0;i<n;i++{
		if len(s) > 0 {
			if s[len(s)-1].cnt == k {
				s = s[:len(s)-k]
			}

			if len(s) > 0 && s[len(s)-1].c == str[i]{
				s = append(s, Pair{str[i], s[len(s)-1].cnt+1})
			}else {
				s = append(s, Pair{str[i], 1})
			}
		}else {
			s = append(s, Pair{str[i], 1})
		}
	}

	if len(s) > 0 && s[len(s)-1].cnt == k {
		s = s[:len(s)-k]
	}

	res := ""
	for i:=0;i<len(s);i++{
		res += string(s[i].c)
	}
	return res
}

func main(){
	fmt.Println(remove_all_duplicates_in_string("deeedbbcccbdaa",3))
}