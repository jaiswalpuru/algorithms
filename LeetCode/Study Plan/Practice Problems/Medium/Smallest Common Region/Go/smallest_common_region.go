package main

import "fmt"

func smallest_region(regions [][]string, region_one, region_two string) string {
	hash_map := make(map[string]string)
	n := len(regions)
	for i := 0; i < n; i++ {
		m := len(regions[i])
		for j := 1; j < m; j++ {
			hash_map[regions[i][j]] = regions[i][0]
		}
	}

	visited := make(map[string]bool)
	common := ""
	for {
		if _, ok := hash_map[region_one]; !ok {
			visited[region_one] = true
			break
		} else {
			visited[region_one] = true
			region_one = hash_map[region_one]
		}
	}

	for {
		if _, ok := hash_map[region_two]; !ok {
			common = region_two
			break
		} else {
			if visited[region_two] {
				common = region_two
				break
			} else {
				region_two = hash_map[region_two]
			}
		}
	}
	return common
}

func main() {
	fmt.Println(smallest_region([][]string{
		{"Earth", "North America", "South America"},
		{"North America", "United States", "Canada"},
		{"United States", "New York", "Boston"},
		{"Canada", "Ontario", "Quebec"},
		{"South America", "Brazil"},
	}, "Quebec", "New York"))
}
