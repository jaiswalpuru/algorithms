package main

func reverse_string(s []byte) {
	n := len(s)

	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

func main() {
	reverse_string([]byte{"h", "e", "l", "l", "o"})
}
