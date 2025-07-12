class TrieNode {
public:
    bool is_word;
    unordered_map<char, TrieNode*> children;
    TrieNode() : is_word(false) {}
};

class WordDictionary {
public:
    TrieNode* root;
    WordDictionary() {
        root = new TrieNode();
    }
    
    void addWord(string word) {
        TrieNode* temp = root;
        for (char c : word) {
            if (temp->children.find(c) == temp->children.end()) temp->children[c] = new TrieNode();
            temp = temp->children[c];
        }
        temp->is_word = true;
    }
    
    bool search(string word) {
        return search_word(word, root);
    }

    bool search_word(string s, TrieNode* temp) {
        for (int i = 0; i < s.size(); i++) {
            char c = s[i];
            if (c == '.') {
                for (auto [k ,v] : temp->children) {
                    if (search_word(s.substr(i + 1), v)) return true;
                }
                return false;
            } else {
                if (temp->children.find(c) == temp->children.end()) return false;
                temp = temp->children[c];
            }
        }
        return temp->is_word;
    }
};

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary* obj = new WordDictionary();
 * obj->addWord(word);
 * bool param_2 = obj->search(word);
 */
