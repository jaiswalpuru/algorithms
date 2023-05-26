class Node {
    Node[] children;
    List<String> word;
    public Node() {
        this.children = new Node[26];
        this.word = new ArrayList<>();
    }
}

class Solution {
    Node trie;

    private void insert(String word) {
        Node root = trie;
        for (int i=0; i<word.length(); i++) {
            int k = word.charAt(i)-'a';
            if (root.children[k] == null) root.children[k] = new Node();
            if (root.children[k].word.size() < 3) root.children[k].word.add(word);
            root = root.children[k];
        }
    }

    private List<List<String>> search(String word) {
        Node root = trie;
        List<List<String>> res = new ArrayList<>();
        for (int i=0; i<word.length(); i++) {
            int k = word.charAt(i)-'a';
            if (root.children[k] == null) root.children[k] = new Node();
            res.add(root.children[k].word);
            root = root.children[k];
        }
        return res;
    }

    public List<List<String>> suggestedProducts(String[] products, String searchWord) {
        Arrays.sort(products);
        trie = new Node();
        for (String prod : products) insert(prod);

        return search(searchWord);
    }
}
