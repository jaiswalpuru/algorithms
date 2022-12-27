package main

import "sort"

type Trie struct {
	str      string
	time     int
	children map[byte]*Trie
}

func New() *Trie { return &Trie{children: make(map[byte]*Trie)} }

func (this *Trie) Insert(word string, cnt int) {
	root := this
	for i := 0; i < len(word); i++ {
		ind := word[i]
		if _, ok := root.children[ind]; !ok {
			root.children[ind] = New()
		}
		root = root.children[ind]
	}
	root.str = word
	root.time += cnt
}

type AutocompleteSystem struct {
	str  string
	trie *Trie
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	t := New()
	for i := 0; i < len(sentences); i++ {
		t.Insert(sentences[i], times[i])
	}
	return AutocompleteSystem{trie: t, str: ""}
}

type Pair struct {
	sen string
	val int
}

func (this *Trie) Search(str string) []string {
	res := []Pair{}
	root := this
	f := true
	for i := 0; i < len(str); i++ {
		char := str[i]
		if _, ok := root.children[char]; !ok {
			f = false
			break
		}
		root = root.children[char]
	}
	if !f {
		return nil
	}
	dfs(root, &res)
	sort.Slice(res, func(i, j int) bool {
		if res[i].val == res[j].val {
			return res[i].sen < res[j].sen
		}
		return res[i].val > res[j].val
	})
	ans := []string{}
	cnt := 0
	for i := 0; i < len(res) && cnt < 3; i++ {
		ans = append(ans, res[i].sen)
		cnt++
	}
	return ans
}

func dfs(trie *Trie, res *[]Pair) {
	if trie.str != "" {
		*res = append(*res, Pair{trie.str, trie.time})
	}
	for _, v := range trie.children {
		dfs(v, res)
	}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.trie.Insert(this.str, 1)
		this.str = ""
		return nil
	}
	this.str += string(c)
	return this.trie.Search(this.str)
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */
