class TrieNode {
    Boolean isWord;
    TrieNode[] children;

    public TrieNode() {
        isWord = false;
        children = new TrieNode[26];
    }
}

class WordDictionary {

    TrieNode root;
    public WordDictionary() {
        root = new TrieNode();
    }
    
    public void addWord(String word) {
        TrieNode root = this.root;
        for (Character c : word.toCharArray()) {
            int ind = c-'a';
            if (root.children[ind] == null) root.children[ind] = new TrieNode();
            root = root.children[ind];
        }
        root.isWord = true;
    }
    
    public boolean search(String word) {
        return searchWord(word, root);
    }

    public boolean searchWord(String word, TrieNode root) {
        StringBuilder sb = new StringBuilder(word);
        for (int i=0; i<sb.length(); i++) {
            if (sb.charAt(i) == '.') {
                for (int j=0; j<root.children.length; j++) {
                    if (root.children[j] != null) {
                        if (searchWord(sb.substring(i+1), root.children[j])) {
                            return true;
                        }
                    }
                }
                return false;
            } else {
                int ind = sb.charAt(i)-'a';
                if (root.children[ind] == null) return false;
                root = root.children[ind];
            }
        }
        return root.isWord;
    }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary obj = new WordDictionary();
 * obj.addWord(word);
 * boolean param_2 = obj.search(word);
 */
