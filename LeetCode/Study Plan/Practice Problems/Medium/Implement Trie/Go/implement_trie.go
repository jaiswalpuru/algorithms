package main

type TrieNode struct {
	is_word bool
	children [26]*TrieNode
}

type Trie struct {
	root *TrieNode    
}


func Constructor() Trie { return Trie{root : &TrieNode{}} }


func (this *Trie) Insert(word string)  {
    root := this.root
    for _, ch := range word {
    	index := ch - 'a'
    	if root.children[index] == nil {
    		root.children[index] = &TrieNode{}
    	}
    	root = root.children[index]
    }
   	root.is_word = true
}


func (this *Trie) Search(word string) bool {
    root := this.root
    for _,ch := range word {
    	index := ch - 'a'
    	if root.children[index] == nil {
    		return false
    	}
    	root = root.children[index]
    }
    return root.is_word
}


func (this *Trie) StartsWith(prefix string) bool {
	root := this.root
	for _, ch := range prefix {
		index := ch - 'a'
		if root.children[index] == nil {
			return false
		}
		root = root.children[index]
	}
	return true
}


func main() {

}