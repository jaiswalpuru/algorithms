package main

import "fmt"

func truncate_sentence(s string, k int) string {
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			k--
			if k == 0 {
				break
			}
		}
		str += string(s[i])
	}
	return str
}

func main() {
	fmt.Println(truncate_sentence("Hello how are you Contestant", 4))
}
