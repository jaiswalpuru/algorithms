package main

import (
	"strconv"
	"strings"
)

type Codec struct{ mp map[int]string }

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	mp := make(map[int]string)
	encoded_string := ""
	n := len(strs)
	for i := 0; i < n; i++ {
		mp[i] = strs[i]
		str := strconv.Itoa(i)
		encoded_string += str + " "
	}
	codec.mp = mp
	return encoded_string[:len(encoded_string)-1]
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	index_string := strings.Split(strs, " ")
	n := len(index_string)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = codec.mp[i]
	}
	return res
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
