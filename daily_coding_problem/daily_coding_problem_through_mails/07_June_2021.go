/*
	Implement a URL shortener with the following methods:
	shorten(url), which shortens the url into a six-character alphanumeric string, such as zLg6wl.
	restore(short), which expands the shortened string into the original url.
	If no such shortened string exists, return null.
	Hint: What if we enter the same URL twice?
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	dict         = make(map[string]string)
	alpha_number = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
)

func create_short_url() string {
	short_url := ""
	for short_url == "" {
		for len(short_url) < 6 {
			short_url += string(alpha_number[rand.Intn(62)])
		}
		if _, ok := dict[short_url]; ok {
			short_url = ""
		}
	}
	return short_url
}

func shorten(url string) string {
	//check in the dictionary if the ur already exists or not
	for k := range dict {
		if dict[k] == url {
			fmt.Println("Shortened url already present : ", url)
			return k
		}
	}
	new_short_url := create_short_url()
	dict[new_short_url] = url
	return new_short_url
}

func restore(encoded string) string {
	if _, ok := dict[encoded]; ok {
		return dict[encoded]
	}
	return ""
}

func main() {

	rand.Seed(time.Now().UnixNano())

	url := []string{"www.facebook.com", "www.google.com", "www.facebook.com"}

	for i := range url {
		shorten_url := shorten(url[i])
		fmt.Println("Shorten url : ", shorten_url, " - Original url : ", restore(shorten_url))
	}
}
