#include <iostream>
#include "trees.h"

using namespace std;

bool dfs(Node* l, Node* r) {
    if (l == nullptr && r == nullptr) return true;
    if (!l || !r) return false;
    if (l->val != r->val) return false;
    return dfs(l->left, r->right) && dfs(l->right, r->left);
}

bool is_symmetric(Node* root) {
    if (root == nullptr) return true;
    return dfs(root->left, root->right);
}

int main(int argc, char** argv) {
    Node* root = new Node(5);
    root->left = new Node(2);
    root->right = new Node(2);
    root->left->left = new Node(1);
    root->right->right = new Node(1);
    cout << is_symmetric(root) << "\n";
    return 0;
}
