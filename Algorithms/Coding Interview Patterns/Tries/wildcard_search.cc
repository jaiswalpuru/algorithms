#include <iostream>
#include "tries.h"

bool _helper(TrieNode* t, int i, string s) {
    for (int j = i; j < s.length(); j++) {
        char c = s[j];
        if (c == '.') {
            for (auto p : t->children) {
                if (_helper(t->children[p.first], j + 1, s)) return true;
            }
            return false;
        } else if (t->children.find(c) != t->children.end()) {
            t = t->children[c];
        } else {
            return false;
        }
    }
    return t->is_word;
}

bool search(Trie* t, string s) {
    return _helper(t->root, 0, s);
}

int main(int argc, char **argv) {
    Trie* t = new Trie();
    t->insert("band");
    t->insert("rat");
    cout << search(t, "ra.") << "\n";
    cout << search(t, "b..") << "\n";
    t->insert("ran");
    cout << search(t, ".an") << "\n";
    return 0;
}
