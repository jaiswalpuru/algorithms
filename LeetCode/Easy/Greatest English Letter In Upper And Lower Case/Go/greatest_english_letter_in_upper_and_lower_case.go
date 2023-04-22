package main

import "fmt"
import "sort"
import "strings"

func greatest_english_letter_in_upper_and_lower_case(str string) string {
	n := len(str)

	if n == 1 {
		return ""
	}

	mp := make(map[string]bool)
	v := []string{}
	for i := 0; i < n; i++ {
		mp[string(str[i])] = true
	}

	for i := 0; i < n; i++ {
		if str[i] >= 'a' && str[i] <= 'z' && mp[strings.ToUpper(string(str[i]))] {
			v = append(v, strings.ToUpper(string(str[i])))
		}
	}

	sort.Strings(v)
	if len(v) == 0 {
		return ""
	}
	return v[len(v)-1]
}

func main() {
	fmt.Println(greatest_english_letter_in_upper_and_lower_case("lEeTcOdE"))
}
