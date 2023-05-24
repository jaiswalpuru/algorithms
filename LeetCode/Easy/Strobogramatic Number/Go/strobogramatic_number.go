package main

import "fmt"

func is_strobogrammtic_number(str string) bool {
	n := len(str)

	i, j := 0, n-1
	for i < j {
		if str[i] == '6' {
			if str[j] != '9' {
				return false
			}
		} else if str[i] == '9' {
			if str[j] != '6' {
				return false
			}
		} else {
			if str[i] != '1' && str[i] != '0' && str[i] != '8' {
				return false
			} else {
				if str[i] != str[j] {
					return false
				}
			}
		}
		i++
		j--
	}
	if i == j {
		if str[i] != '1' && str[i] != '0' && str[i] != '8' {
			return false
		}
	}
	return true
}

//-------------------------A better approach coding wise not complexity wise(which is same as above)------------------
func is_strobogrammtic_number_better(str string) bool {
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

//--------------------------------------------------------------------------------------------------------------------

func main() {
	fmt.Println(is_strobogrammtic_number_better("69"))
}