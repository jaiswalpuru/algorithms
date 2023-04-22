package main

import "fmt"

func move_pieces_to_obtain_a_string(start, target string) bool {
	s, e := []byte(start), []byte(target)
	n := len(s)
	i, j := 0, 0
	for i < n || j < n {
		for i < n && s[i] == '_' {
			i++
		}
		for j < n && e[j] == '_' {
			j++
		}
		if i == n || j == n || (s[i] != e[j]) || (s[i] == 'L' && i < j) || (s[i] == 'R' && i > j) {
			break
		}
		i++
		j++
	}
	return i == n && j == n
}

func main() {
	fmt.Println(move_pieces_to_obtain_a_string("_L__R__R_", "L______RR"))
}
