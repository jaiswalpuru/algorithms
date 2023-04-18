class TrieNode {
    boolean isWord;
    TrieNode[] children;
    public TrieNode() {
        children = new TrieNode[26];
    }
}

class Solution {
    
    public void insert(TrieNode root, String word) {
        TrieNode temp = root;
        for (char c : word.toCharArray()) {
            int index = c - 'a';
            if (temp.children[index] == null) {
                temp.children[index] = new TrieNode();
            }
            temp = temp.children[index];
        }
        temp.isWord = true;
    }

    public boolean find(TrieNode root, String word) {
        TrieNode temp = root;
        for (char c : word.toCharArray()) {
            int index = c - 'a';
            if (!temp.children[index].isWord) {
                return false;
            }
            temp = temp.children[index];
        }
        return true;
    } 

    public String longestWord(String[] words) {
        Arrays.sort(words);
        TrieNode root = new TrieNode();
        String ans = "";
        for (String word : words) insert(root, word);
        for (int i=words.length-1; i>=0; i--) {
            if (find(root, words[i])) {
                if (ans.length() < words[i].length()) {
                    ans = words[i];
                } else {
                    if (ans.length() == words[i].length()) ans = words[i];
                }
            }
        }
        return ans;
    }
}
