package main

type MagicDictionary struct {
	mp            map[string]bool
	string_length map[int]bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{mp: make(map[string]bool), string_length: make(map[int]bool)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := 0; i < len(dictionary); i++ {
		this.mp[dictionary[i]] = true
		this.string_length[len(dictionary[i])] = true
	}
}

func (this *MagicDictionary) Search(word string) bool {
	if !this.string_length[len(word)] {
		return false
	}
	st := []byte(word)
	for i := 0; i < len(st); i++ {
		t := st[i]
		for j := 0; j < 26; j++ {
			st[i] = byte(int('a') + j)
			if t == st[i] {
				continue
			}
			if this.mp[string(st)] {
				return true
			}
		}
		st[i] = t
	}
	return false
}
