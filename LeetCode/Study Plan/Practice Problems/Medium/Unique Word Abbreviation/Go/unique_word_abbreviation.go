type ValidWordAbbr struct {
	hash_map map[string]map[string]bool
}

func Constructor(dictionary []string) ValidWordAbbr {
	hash_map := make(map[string]map[string]bool)
	for i := 0; i < len(dictionary); i++ {
		word := dictionary[i]
		word = string(word[0]) + strconv.Itoa(len(word)-2) + string(word[len(word)-1])
		if _, ok := hash_map[word]; !ok {
			hash_map[word] = make(map[string]bool)
		}
		hash_map[word][dictionary[i]] = true
	}
	return ValidWordAbbr{hash_map: hash_map}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	str := string(word[0]) + strconv.Itoa(len(word)-2) + string(word[len(word)-1])
	if v, ok := this.hash_map[str]; ok {
		fmt.Println(v)
		if v[word] {
			return len(v) == 1
		} else {
			return false
		}
	}
	return true
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */
