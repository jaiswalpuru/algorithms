package main

import "fmt"

var ascii_lower = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
	'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

type Pair struct {
	word string
	path []string
}

func word_ladder(start, end string, words []string) []string {
	queue := []Pair{}
	p := []string{start}
	queue = append(queue, Pair{word: start, path: p})

	for len(queue) > 0 {
		fmt.Println("Before popping : ", queue)
		word, path := queue[0].word, queue[0].path
		queue = queue[1:]
		fmt.Println("After pooping : ", queue)
		if word == end {
			return path
		}

		mp := make(map[string]int)
		for i := 0; i < len(words); i++ {
			mp[words[i]] = i
		}

		for i := 0; i < len(word); i++ {
			for _, v := range ascii_lower {
				nxt_word := word[:i] + string(v) + word[i+1:]
				if _, ok := mp[nxt_word]; ok {
					v := mp[nxt_word]
					temp := words[:v]
					temp = append(temp, words[v+1:]...)
					words = temp
					tmp := path
					tmp = append(tmp, nxt_word)
					queue = append(queue, Pair{word: nxt_word, path: tmp})
				}
			}
		}
	}

	return nil
}

func main() {
	start := "dog"
	end := "cat"
	words := []string{"dot", "dop", "dat", "cat"}
	fmt.Println(word_ladder(start, end, words))
}
