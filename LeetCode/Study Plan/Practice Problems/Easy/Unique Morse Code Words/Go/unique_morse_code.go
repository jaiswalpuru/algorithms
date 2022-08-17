package main

import "fmt"

var c = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func unique_morse_code_representations(words []string) int {
	n := len(words)
	mp := make(map[string]struct{})
	for i := 0; i < n; i++ {
		t := ""
		for j := 0; j < len(words[i]); j++ {
			t += c[int(words[i][j]-'a')]
		}
		mp[t] = struct{}{}
	}
	return len(mp)
}

func main() {
	fmt.Println(unique_morse_code_representations([]string{"gin", "zen", "gig", "msg"}))
}
