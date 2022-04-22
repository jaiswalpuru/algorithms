package main

type Trie struct {
	val      int
	hash_map map[string]*Trie
}

type FileSystem struct {
	root *Trie
}

func Constructor() FileSystem { return FileSystem{root: &Trie{hash_map: make(map[string]*Trie)}} }

func (this *FileSystem) CreatePath(path string, value int) bool {
	str := strings.Split(path, "/")
	dir := this.root
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] != "" {
			if i != n-1 {
				if _, ok := dir.hash_map[str[i]]; !ok {
					return false
				}
				dir = dir.hash_map[str[i]]
			} else {
				if _, ok := dir.hash_map[str[i]]; ok {
					return false
				}
				dir.hash_map[str[i]] = &Trie{val: value, hash_map: make(map[string]*Trie)}
			}
		}
	}
	return true
}

func (this *FileSystem) Get(path string) int {
	str := strings.Split(path, "/")
	dir := this.root
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] != "" {
			if _, ok := dir.hash_map[str[i]]; !ok {
				return -1
			}
			dir = dir.hash_map[str[i]]
		}
	}
	return dir.val
}

func main() {}
