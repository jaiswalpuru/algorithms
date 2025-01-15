#include <iostream>
#include "tries.h"

int main(int argc, char **argv) {
    Trie* t = new Trie();
    t->insert("top");
    t->insert("bye");
    cout << t->has_prefix("to") << "\n";
    cout << t->search("to") << "\n";
    t->insert("to");
    cout << t->search("to") << "\n";
    return 0;
}
