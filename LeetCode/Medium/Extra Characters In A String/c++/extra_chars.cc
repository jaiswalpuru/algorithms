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
        Node *temp = node;

        for (char c : word) {
            if (temp->children.find(c) == temp->children.end()) temp->children[c] = new Node();
            temp = temp->children[c];
        }
        temp->is_word = true;
    }

    int search(string word, int i, unordered_map<int, int> &memo) {
        if (i == word.length()) return 0;
        Node *temp = node;
        
        if (memo.count(i)) return memo[i];
        int res = search(word, i+1, memo) + 1;
        for (int j = i; j < word.length(); j++) {
            char c = word[j];
            if (temp->children.find(c) == temp->children.end()) break;
            temp = temp->children[c];
            if (temp->is_word) res = min(res, search(word, j+1, memo));
        }
        return memo[i] = res;
    }
};

class Solution {
public:
    int minExtraChar(string s, vector<string>& dictionary) {
        Trie *t = new Trie();
        unordered_map<int, int> memo;
        for(string s : dictionary) t->insert(s);

        return t->search(s, 0, memo);
    }
};
