package main

import "fmt"

func find_unique_binary_string(nums []string) string {
	hash_map := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		hash_map[nums[i]] = true
	}
	res := ""
	back_track([]byte(nums[0]), hash_map, &res, 0)
	return res
}

func back_track(str []byte, hash_map map[string]bool, res *string, ind int) bool {
	if !hash_map[string(str)] {
		*res = string(str)
		return true
	}
	for i := ind; i < len(str); i++ {
		if str[i] == '0' {
			str[i] = '1'
			if back_track(str, hash_map, res, i+1) {
				return true
			}
			str[i] = '0'
		} else {
			str[i] = '0'
			if back_track(str, hash_map, res, i+1) {
				return true
			}
			str[i] = '1'
		}
	}
	return false
}

func main() {
	fmt.Println(find_unique_binary_string([]string{"01", "10"}))
}
