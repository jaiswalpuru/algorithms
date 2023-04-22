class TrieNode {
    TrieNode[] children;
    Boolean isWord;
    public TrieNode() {
        children = new TrieNode[26];
        isWord = false;
    }
}

class Trie {
    
    TrieNode root;
    public Trie() {
        root = new TrieNode();
    }
    
    public void insert(String word) {
        TrieNode r = this.root;
        for (char c : word.toCharArray()) {
            if (r.children[c-'a'] == null) {
                r.children[c-'a'] = new TrieNode();
            }
            r = r.children[c-'a'];
        }
        r.isWord = true;
    }
    
    public boolean search(String word) {
        TrieNode r = this.root;
        for (char c : word.toCharArray()) {
            if (r.children[c-'a'] == null) {
                return false;
            }
            r = r.children[c-'a'];
        }
        return r.isWord ? true : false;
    }
    
    public boolean startsWith(String prefix) {
        TrieNode r = this.root;
        for (char c : prefix.toCharArray()) {
            if (r.children[c-'a'] == null) {
                return false;
            }
            r = r.children[c-'a'];
        }
        return true;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */
