package main

import "fmt"

func letter_tile_possibilities(s string) int {
	visited_1 := make(map[string]bool)
	visited_2 := make(map[int]bool)
	set := []string{}
	_letter_tile_possibilities(s, &set, "", &visited_1, &visited_2)
	return len(set)
}

func _letter_tile_possibilities(s string, set *[]string, t string, v1 *map[string]bool, v2 *map[int]bool) {
	for i := 0; i < len(s); i++ {
		if !(*v2)[i] {
			(*v2)[i] = true
			temp := t + string(s[i])
			if !(*v1)[temp] {
				(*v1)[temp] = true
				*set = append(*set, temp)
			}
			_letter_tile_possibilities(s, set, temp, v1, v2)
			(*v2)[i] = false
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(letter_tile_possibilities("AAB"))
}
