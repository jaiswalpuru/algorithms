package main

import "fmt"

// this idea is based on grouping the names with the first character
// then finiding the mutual among them.
func distinctNames(ideas []string) int64 {
	group := make(map[int]map[string]bool)
	size := len(ideas)

	//grouping by character
	for i := 0; i < size; i++ {
		firstChar := int(ideas[i][0] - 'a')
		if _, ok := group[firstChar]; !ok {
			group[firstChar] = make(map[string]bool)
		}
		group[firstChar][ideas[i][1:]] = true
	}

	res := int64(0)
	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			mutual := int64(0)
			for k, _ := range group[i] {
				if _, ok := group[j]; ok {
					if group[j][k] {
						mutual++
					}
				}
			}
			res += 2 * (int64(len(group[i])) - mutual) * (int64(len(group[j])) - mutual)
		}
	}

	return res
}

// *** BRUTE FORCE ***
// this is brute force method which anyone can come up with
func distinctNamesBruteForce(ideas []string) int64 {
	companyNames := make(map[string]bool)
	size := len(ideas)

	for i := 0; i < size; i++ {
		companyNames[ideas[i]] = true
	}

	res := int64(0)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			s1 := []byte(ideas[i])
			s2 := []byte(ideas[j])
			s1[0], s2[0] = s2[0], s1[0]
			if !companyNames[string(s1)] && !companyNames[string(s2)] {
				res += 2
			}
		}
	}
	return res
}

func main() {
	fmt.Println(distinctNames([]string{"coffee", "donuts", "time", "toffee"}))
}
