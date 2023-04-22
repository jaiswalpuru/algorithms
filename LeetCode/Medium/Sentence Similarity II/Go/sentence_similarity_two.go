package main

import "fmt"

func sentence_similarity_two(sen_1 []string, sen_2 []string, sim_pair [][]string) bool {
	if len(sen_1) != len(sen_2) {
		return false
	}

	mp := make(map[string][]string)
	n := len(sim_pair)

	for i := 0; i < n; i++ {
		mp[sim_pair[i][0]] = append(mp[sim_pair[i][0]], sim_pair[i][1])
		mp[sim_pair[i][1]] = append(mp[sim_pair[i][1]], sim_pair[i][0])
	}

	for i := 0; i < len(sen_1); i++ {
		if !is_similar(sen_1[i], sen_2[i], mp, map[string]bool{}) {
			return false
		}
	}
	return true
}

func is_similar(a, b string, mp map[string][]string, visited map[string]bool) bool {
	if a == b {
		return true
	}
	visited[a] = true
	v := false
	for i := 0; i < len(mp[a]); i++ {
		if !visited[mp[a][i]] {
			v = v || is_similar(mp[a][i], b, mp, visited)
		}
	}
	return v
}

func main() {
	fmt.Println(sentence_similarity_two([]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, [][]string{
		{"great", "good"}, {"fine", "good"}, {"drama", "acting"}, {"skills", "talent"},
	}))
}
