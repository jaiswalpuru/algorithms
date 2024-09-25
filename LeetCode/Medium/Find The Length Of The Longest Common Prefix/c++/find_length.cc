class Node {
public:
    vector<Node*> children;
    char ch;
    Node() : children(10, nullptr) {}
    ~Node() {
        for (Node* child : children) { delete child; }
    }
};

class Trie {
public:
    Node* node;
    Trie() : node{new Node()} {}
    
    void insert(string word) {
        int ind = 0;
        Node *temp = node;
        for (int i = 0; i < word.length(); i++) {
            ind = word[i]-'0';
            if (temp->children[ind] == nullptr) temp->children[ind] = new Node();
            temp = temp->children[ind]; 
            temp->ch = word[i];
        }
    }

    int search(string word) {
        int len = 0;
        int ind = 0;
        Node *temp = node;
        for (int i = 0; i < word.length(); i++) {
            ind = word[i] - '0';
            if (temp->children[ind] == nullptr) break;
            temp = temp->children[ind];
            if (temp->ch == word[i]) len++;
        }
        return len;
    }

    ~Trie() { delete node; }
};

class Solution {
    Trie* trie;
public:
    int longestCommonPrefix(vector<int>& arr1, vector<int>& arr2) {
        int res = 0;
        trie = new Trie();
        for (int i = 0; i < arr1.size(); i++) {
            trie->insert(to_string(arr1[i]));
        }

        for (int i = 0; i < arr2.size(); i++) {
            res = max(res, trie->search(to_string(arr2[i])));
        }

        return res;
    }
};
