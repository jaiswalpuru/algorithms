package main

import "fmt"

func rearrange_spaces_between_words(text string) string {
	sp := 0
	words := []string{}
	str := ""
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			sp++
			if str != "" {
				words = append(words, str)
				str = ""
			}
		} else {
			str += string(text[i])
		}
	}
	if str != "" {
		words = append(words, str)
	}
	str = ""
	spaces := 0
	if len(words) == 1 {
		spaces = sp
	} else {
		spaces = sp / (len(words) - 1)
	}
	for i := 0; i < len(words); i++ {
		str += words[i]
		if i == len(words)-1 {
			break
		}
		for j := 0; j < spaces; j++ {
			str += " "
		}
	}
	extra_space := sp - (spaces * (len(words) - 1))
	for i := 0; i < extra_space; i++ {
		str += " "
	}
	return str
}

func main() {
	fmt.Println(rearrange_spaces_between_words("  this   is  a sentence "))
}
