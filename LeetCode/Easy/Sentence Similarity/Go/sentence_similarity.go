package main

import "fmt"

func sentence_similarity(sen_1 []string, sen_2 []string, sim [][]string) bool {
    if len(sen_1) != len(sen_2) {
		return false
	}

	mp := make(map[string][]string)
	for i := 0; i < len(sim); i++ {
        mp[sim[i][0]] = append(mp[sim[i][0]], sim[i][1])
        mp[sim[i][1]] = append(mp[sim[i][1]], sim[i][0])
	}

	for i := 0; i < len(sen_1); i++ {
		a, b := sen_1[i], sen_2[i]
        if a == b {
            continue
        }
        l := mp[a]
        found := false
        for j:=0; j<len(l); j++ {
            if l[j] == b {
                found = true
                break
            }
        }
        if !found{
            return false
        }
	}

	return true
}

func main() {
	fmt.Println(sentence_similarity([]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, [][]string{
		{"great", "fine"}, {"drama", "acting"}, {"skills", "talent"},
	}))
}
