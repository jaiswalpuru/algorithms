#include <iostream>
#include "trees.h"

using namespace std;

void dfs(Node *root) {
    if (root == nullptr) return;
    cout << root->val << " ";
    dfs(root->left);
    dfs(root->right);
}

int main(int argc, char **argv) {
    Node* root = new Node;
    root->val = 1;
    root->left = new Node;
    root->left->val = 2;
    root->right = new Node;
    root->right->val = 3;
    dfs(root);
    cout << "\n";
    return 0;
}
