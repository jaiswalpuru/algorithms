package main

import (
	"fmt"
)

func is_strobogrammtic_number_better(str string) bool {
	if str[0] == '0' {
		return false
	}
	rotated_string := []byte{}
	n := len(str)
	for i := n - 1; i >= 0; i-- {
		if str[i] == '0' || str[i] == '1' || str[i] == '8' {
			rotated_string = append(rotated_string, str[i])
		} else if str[i] == '6' {
			rotated_string = append(rotated_string, '9')
		} else if str[i] == '9' {
			rotated_string = append(rotated_string, '6')
		} else {
			return false
		}
	}
	return string(rotated_string) == str
}

//----------------------------Brute force DO NOT USE THIS---------------------------------(TLE)
func generate_all_strobogramatic_numbers(n int) []string {
	if n == 1 {
		return []string{"1", "0", "8"}
	}
	strobo := []string{"1", "0", "8", "6", "9"}
	res := []string{}
	_generate_all_strobogramatic_numbers(n, &res, "", strobo, 0)
	return res
}

func _generate_all_strobogramatic_numbers(n int, res *[]string, set string, strobo []string, ind int) {
	if len(set) == n {
		if is_strobogrammtic_number_better(set) {
			*res = append(*res, set)
		}
		return
	}

	for i := ind; i < len(strobo); i++ {
		temp := set + strobo[i]
		_generate_all_strobogramatic_numbers(n, res, temp, strobo, ind)
		temp = temp[:len(temp)-1]
	}
}

//----------------------------DO NOT USE THIS---------------------------------

//----------------------------Efficient solution---------------------------------
func generate_all_strobogramatic_numbers_eff(n int) []string {
	reversible_pairs := [][]byte{{'1', '1'}, {'0', '0'}, {'6', '9'}, {'9', '6'}, {'8', '8'}}
	return _generate_all_strobogramatic_numbers_eff(n, n, reversible_pairs)
}

func _generate_all_strobogramatic_numbers_eff(n, l int, reversible_pairs [][]byte) []string {
	if n == 0 {
		return []string{""}
	}

	if n == 1 {
		return []string{"1", "8", "0"}
	}

	curr_res := []string{}
	prev_strob_num := _generate_all_strobogramatic_numbers_eff(n-2, l, reversible_pairs)
	for i := 0; i < len(prev_strob_num); i++ {
		for j := 0; j < len(reversible_pairs); j++ {
			if reversible_pairs[j][0] != '0' || n != l {
				curr_res = append(curr_res, string(reversible_pairs[j][0])+prev_strob_num[i]+string(reversible_pairs[j][1]))
			}
		}
	}
	return curr_res
}

//----------------------------Efficient-----------------------------------------------

func main() {
	fmt.Println(generate_all_strobogramatic_numbers_eff(3))
}
