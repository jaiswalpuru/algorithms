package main

import "fmt"

func get_hash_key(str string) string {
	str_r := []rune(str)
	n := len(str_r)
	shift := str_r[0]
	for i := 0; i < n; i++ {
		str_r[i] = (str_r[i]-shift+26)%26 + 'a'
	}

	return string(str_r)
}

func group_shifted_strings(arr []string) [][]string {
	hash_map := make(map[string][]string)
	n := len(arr)
	for i := 0; i < n; i++ {
		hash_map[get_hash_key(arr[i])] = append(hash_map[get_hash_key(arr[i])], arr[i])
	}
	res := [][]string{}
	for _, v := range hash_map {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(group_shifted_strings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
}
