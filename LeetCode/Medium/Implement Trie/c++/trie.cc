class TrieNode {
public:
    bool is_word;
    unordered_map<char, TrieNode*> children;
    TrieNode() : is_word(false) {} 
};

class Trie {
public:
    TrieNode* root;
    Trie() {
        root = new TrieNode();
    }
    
    void insert(string word) {
        TrieNode* temp = this->root;
        for (char c : word) {
            if (temp->children.find(c) == temp->children.end()) {
                temp->children[c] = new TrieNode();
            }
            temp = temp->children[c];
        }
        temp->is_word = true;
    }
    
    bool search(string word) {
        TrieNode* temp = this->root;
        for (char c : word) {
            if (temp->children.find(c) == temp->children.end()) return false;
            temp = temp->children[c];
        }
        return temp->is_word;
    }
    
    bool startsWith(string prefix) {
        TrieNode* temp = root;
        for (char c : prefix) {
            if (temp->children.find(c) == temp->children.end()) return false;
            temp = temp->children[c];
        }
        return temp->children.size() > 0 || temp->is_word;
    }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */
