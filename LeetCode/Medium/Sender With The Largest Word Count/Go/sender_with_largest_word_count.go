package main

import "fmt"
import "strings"

func sender_with_largest_word_count(mes []string, senders []string) string {
	mp := make(map[string]int)

	res := ""
	max_cnt := 0
	for i := 0; i < len(mes); i++ {
		mp[senders[i]] += len(strings.Split(mes[i], " "))
		if max_cnt <= mp[senders[i]] {
			if max_cnt == mp[senders[i]] {
				if senders[i] > res {
					res = senders[i]
				}
			} else {
				max_cnt = mp[senders[i]]
				res = senders[i]
			}
		}
	}

	return res
}

func main() {
	fmt.Println(sender_with_largest_word_count([]string{
		"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree",
	}, []string{
		"Alice", "userTwo", "userThree", "Alice",
	}))
}
