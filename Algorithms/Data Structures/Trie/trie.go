package trie

//Trie -> datastructure type
type Trie struct {
	letter   rune
	children []*Trie
	meta     map[string]interface{}
	isLeaf   bool
}

//Inserting a node in trie
/*
1. set a current node as root node
2. set the current letter as the first letter of the word
3. if current node has reference to the current letter then set the current node to that referenced node else create a new node set the letter equal to the current letter and also initialize current node to this new node
4. repeat step 3 until the key is traversed
*/

func (t *Trie) hasChild(a rune) (bool, *Trie) {
	for _, child := range t.children {
		if child.letter == a {
			return true, child
		}
	}
	return false, nil
}

//NewTrie creates a new trie with default values
func NewTrie() *Trie {
	nT := &Trie{}
	nT.children = []*Trie{}
	nT.meta = make(map[string]interface{})
	return nT
}

func (t *Trie) addChild(a rune) *Trie {
	nw := NewTrie()
	nw.letter = a
	t.children = append(t.children, nw)
	return nw
}

//Add -> add word in trie
func (t *Trie) Add(word string) *Trie {
	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			node = node.addChild(letters[i])
		}
		i++
		if i == n {
			node.isLeaf = true
		}
	}
	return node
}

//FindNode -> searches whether a node is present in the trie or not
func (t *Trie) FindNode(word string) *Trie {
	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			return nil
		}
		i++
	}
	return node
}

//Find -> wrapper for FindNode
func (t *Trie) Find(word string) *Trie {
	node := t.FindNode(word)

	if node == nil {
		return nil
	}
	if node.isLeaf != true {
		return nil
	}
	return node
}

//Remove a word from the tree
/*
1. check whether the element is already a part of the tree
2. if the element is found then remove it
*/
func (t *Trie) Remove(word string) {
	a := t.Find(word)
	if a != nil {
		a.isLeaf = false
	}
}

//Counting the number of strings in trie
func (t *Trie) Count() int {
	count := 0
	for _, child := range t.children {
		if child.isLeaf == true {
			count++
		}
		count += child.Count()
	}
	return count
}

//Setting / Getting metadata

//Get -> to get the metadata for the key
func (t *Trie) Get(key string) (interface{}, bool) {
	if t == nil {
		return nil, false
	}

	if _, ok := t.meta[key]; ok {
		return t.meta[key], true
	}
	return nil, false
}

//Set -> to set the meta data
func (t *Trie) Set(key string, val interface{}) {
	if t == nil {
		return
	}
	t.meta[key] = val
}
