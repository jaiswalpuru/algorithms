package main

import ("fmt"; "index/suffixarray")

//---------------------do not follow this approach---------------------------- (using suffix array yet to come)
//Idea is to generate all the substrings and dump them into a set based data structure
func number_of_distinct_substrings_in_string_brute_force(str string) int {
	hash_map := make(map[string]string)
	n := len(str)

	for i:=0;i<n;i++{
		t := string(str[i])
		hash_map[t]=t
		for j:=i+1; j<n;j++{
			t += string(str[j])
			hash_map[t] = t
		}
	}
	return len(hash_map)
}
//---------------------do not follow this approach----------------------------

func main() {
	fmt.Println(number_of_distinct_substrings_in_string_brute_force("aabbaba"))
}