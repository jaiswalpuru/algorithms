package main

import "fmt"

//---------------------method 1--------------------------
func permutations_without_dups(s string) []string {
	res := []string{}
	sb := []byte(s)
	mp := make(map[byte]bool)
	recur(sb, &mp, []byte{}, &res)
	return res
}

func recur(sb []byte, mp *map[byte]bool, set []byte, res *[]string) {
	if len(set) == len(sb) {
		*res = append(*res, string(set))
		return
	}
	for i := 0; i < len(sb); i++ {
		if !(*mp)[sb[i]] {
			t := append(set, sb[i])
			(*mp)[sb[i]] = true
			recur(sb, mp, t, res)
			(*mp)[sb[i]] = false
			t = t[:len(t)-1]
		}
	}
}

//---------------------method 1--------------------------

//---------------------method 2--------------------------
func permutations_without_dups_m2(s string) []string {
	sb := []byte(s)
	res := []string{}
	_recur(sb, &res, 0)
	return res
}

func _recur(sb []byte, res *[]string, ind int) {
	if ind == len(sb) {
		*res = append(*res, string(sb))
		return
	}
	for i := ind; i < len(sb); i++ {
		sb[i], sb[ind] = sb[ind], sb[i]
		_recur(sb, res, ind+1)
		sb[i], sb[ind] = sb[ind], sb[i]
	}
}

//---------------------method 2--------------------------

func main() {
	fmt.Println(permutations_without_dups_m2("abc"))
}
