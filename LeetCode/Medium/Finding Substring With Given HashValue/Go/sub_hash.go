package main

import "fmt"

func find_sub_with_hash(str string, power int, modulo int, k int, hash_val int) string {
	hash := 0

	n := len(str)
	ind := -1
	pr := 1
	for i := n - 1; i >= 0; i-- {
		hash = ((hash * power) + int(str[i]-'a'+1)) % modulo

		if i+k >= n {
			pr = (pr * power) % modulo
		} else {
			hash = (hash - (int(str[i+k]-'a'+1)*pr)%modulo + modulo) % modulo
		}

		if hash == hash_val {
			ind = i
		}
	}

	return str[ind : ind+k]
}

func main() {
	fmt.Println(find_sub_with_hash("leetcode", 7, 20, 2, 0))
}
