package main

type TrieNode struct {
	children []*TrieNode
	wt       int
}

func New() *TrieNode { return &TrieNode{children: make([]*TrieNode, 27), wt: 0} }

type WordFilter struct{ root *TrieNode }

func Constructor(words []string) WordFilter {
	t := New()
	n := len(words)

	for i := 0; i < n; i++ {
		w := "{" + words[i]
		insert(t, w, i)
		for j := 0; j < len(w); j++ {
			insert(t, w[j+1:]+w, i)
		}
	}
	return WordFilter{t}
}

func insert(root *TrieNode, word string, wt int) {
	r := root
	for _, c := range word {
		ch := c - 'a'
		if r.children[ch] == nil {
			r.children[ch] = New()
		}
		r = r.children[ch]
		r.wt = wt
	}
}

func (this *WordFilter) F(pre, suff string) int {
	str := suff + "{" + pre
	r := this.root
	for _, c := range str {
		if r.children[c-'a'] == nil {
			return -1
		}
		r = r.children[c-'a']
	}
	return r.wt
}

func main() {}
