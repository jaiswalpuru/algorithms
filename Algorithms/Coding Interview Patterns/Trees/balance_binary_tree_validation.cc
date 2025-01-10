#include <iostream>
#include "trees.h"

using namespace std;

int is_balanced(Node*);

int main(int argc, char **argv) {
    Node* root = new Node;
    root->val = 5;
    root->left = new Node;
    root->left->val = 2;
    root->left->left = new Node;
    root->left->left->val = 1;
    root->left->right = new Node;
    root->left->right->val = 4;
    root->left->right->left = new Node;
    root->left->right->left->val = 3;
    root->right = new Node;
    root->right->val = 7;
    root->right->right = new Node;
    root->right->right->val = 9;
    root->right->right->left = new Node;
    root->right->right->left->val = 6;
    cout << (is_balanced(root) == -1) << "\n";
    return 0;
}

int is_balanced(Node* root) {
    if (!root) return 0;

    int l = is_balanced(root->left);
    int r = is_balanced(root->right);
    if (l == -1 || r == -1) return -1;

    if (abs(l - r) > 1) return -1;
    return 1 + max(l, r);
}

