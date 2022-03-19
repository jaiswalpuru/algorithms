package main

import (
	"fmt"
	"strconv"
)

//----------------------------using recursion-------------------------------------
func different_ways_to_add_paranthesis(str string) []int {
	res := []int{}
	for i:=0; i<len(str); i++{
		if str[i]=='-' || str[i] == '+' || str[i] == '*' {
			left := str[0:i]
			right := str[i+1:len(str)]
			left_part := different_ways_to_add_paranthesis(left)
			right_part := different_ways_to_add_paranthesis(right)
			for j:=0; j<len(left_part); j++{
				for k:=0; k<len(right_part); k++{
					val := 0
					switch str[i] {
					case '+' : 
						val = left_part[j] + right_part[k]
					case '-' :
						val = left_part[j] - right_part[k]
					case '*' : 
						val = left_part[j] * right_part[k]
					}
					res = append(res, val)
				}
			}
		}
	}
	if len(res) == 0 {
		val, _ := strconv.Atoi(str)
		res = append(res, val)
	}
	return res
}
//----------------------------------------------------------------------------------


//---------------------------using memoization--------------------------------------
func different_ways_to_add_paranthesis_memo(str string) []int{
	hash_map := make(map[string][]int)
	return _different_ways_to_add_paranthesis_memo(str, &hash_map)
}

func _different_ways_to_add_paranthesis_memo(str string, hash_map *map[string][]int) []int {
	if _, ok := (*hash_map)[str]; ok {
		return (*hash_map)[str]
	}
	res := []int{}
	for i:=0; i<len(str); i++ {
		if str[i] == '+' || str[i] == '-' || str[i] == '*' {
			left := str[0:i]
			right := str[i+1:]
			left_part := _different_ways_to_add_paranthesis_memo(left, hash_map)
			right_part := _different_ways_to_add_paranthesis_memo(right, hash_map)
			for j:=0; j<len(left_part); j++ {
				for k:=0; k<len(right_part); k++ {
					val := 0
					switch str[i] {
					case '+' : 
						val = left_part[j] + right_part[k]
					case '-' :
						val = left_part[j] - right_part[k]
					case '*' : 
						val = left_part[j] * right_part[k]
					}
					res = append(res, val)
				}
			}
		}
	}

	if len(res) == 0{
		val, _:= strconv.Atoi(str)
		res = append(res, val)
	}
	(*hash_map)[str] = res
	return res
}

//----------------------------------------------------------------------------------
func main() {
	fmt.Println(different_ways_to_add_paranthesis_memo("2-1-1"))
}