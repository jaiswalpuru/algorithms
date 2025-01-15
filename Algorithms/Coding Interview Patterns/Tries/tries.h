#include <string>
#include <unordered_map>

using namespace std;

typedef struct TrieNode {
    bool is_word;
    string word;
    unordered_map<char, struct TrieNode*> children;
} Node;

class Trie {
public:
    Node *root;
    Trie() {
        root = new Node;
    }

    void insert(string s) {
        Node* temp = this->root;
        for (char c : s) {
            if (temp->children.find(c) == temp->children.end()) temp->children[c] = new Node;
            temp = temp->children[c];
        }
        temp->is_word = true;
        temp->word = s;
    }

    bool search(string s) {
        Node* temp = this->root;
        for (char c : s) {
            if (temp->children.find(c) == temp->children.end()) return false;
            temp = temp->children[c];
        }
        return temp->is_word;
    }

    bool has_prefix(string s) {
        Node* temp = this->root;
        for (char c : s) {
            if (temp->children.find(c) == temp->children.end()) return false;
            temp = temp->children[c];
        }
        return true;
    }
};
