class Node {
public:
    unordered_map<char, Node*> children;
    bool is_word;
    Node() : is_word(false) {}
};

class Trie {
public:
    Node *node;
    Trie() : node(new Node()) {}
    
    void insert(string word) {
        Node *t = node;
        for (char c : word) {
            if (t->children.find(c) == t->children.end()) t->children[c] = new Node();
            t = t->children[c];
        }
        t->is_word = true;
    }
    
    bool search(string word) {
        Node *t = node;
        for (char c : word) {
            if (t->children.find(c) == t->children.end()) return false;
            t = t->children[c];
        }
        return t->is_word;
    }
    
    bool startsWith(string prefix) {
        Node *t = node;
        int i = 0;
        for (; i < prefix.length(); i++) {
            if (t->children.find(prefix[i]) == t->children.end()) break;
            t = t->children[prefix[i]];
        }
        return i == prefix.length();
    }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */
