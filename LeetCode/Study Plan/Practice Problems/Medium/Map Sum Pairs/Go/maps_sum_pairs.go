type TrieNode struct {
	val int
	is_word bool
	children [26]*TrieNode
}

type MapSum struct { root *TrieNode }

func Constructor() MapSum {return MapSum{root : &TrieNode{}}}

func (this *MapSum) Insert(key string, val int) {
	root := this.root
	for _, ch := range key {
		index := ch-'a'
		if root.children[index] == nil {
			root.children[index] = &TrieNode{}
		}
		root = root.children[index]
	}
	root.val = val
	root.is_word = true
}

func (this *MapSum) Sum(prefix string) int {
	root := this.root
	for _, ch := range prefix {
		index := ch-'a'
		if root.children[index] == nil {
			return 0
		}
		root = root.children[index]
	}

	sum := 0
	children := root.children
	if root.is_word {
		sum += root.val
	}
    
	for i:=0;i<len(children);i++{
        if children[i] != nil {
            get_sum(children[i], &sum)   
        }
	}
	return sum
}

func get_sum(child *TrieNode, sum *int) {
    if child == nil {
        return
    }
    if child.is_word {
        *sum += child.val
    }
	for i:=0; i<len(child.children); i++{
		if child.children[i] != nil {
			get_sum(child.children[i], sum)
		}
	}
}