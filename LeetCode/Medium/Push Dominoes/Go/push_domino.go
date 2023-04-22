package main

import "fmt"

func push_domino(str string) string {
	str = "L" + str + "R"
	b := []byte(str)
	j := 0
	n := len(b)
	for i := 0; i < n; {
		if b[i] != '.' {
			i++
			j++
		} else {
			for j < n && b[j] == '.' {
				j++
			}
			temp := j
			i--
			if b[i] == 'R' && b[j] == 'L' {
				for i < j {
					b[i] = 'R'
					b[j] = 'L'
					i++
					j--
				}
			} else if b[i] == b[j] {
				char := b[i]
				for i < j {
					b[i] = char
					i++
				}
			}
			i = temp + 1
			j = temp + 1
		}
	}
	b = b[1:]
	b = b[:len(b)-1]
	return string(b)
}

func main() {
	fmt.Println(push_domino(".L.R...LR..L.."))
}
