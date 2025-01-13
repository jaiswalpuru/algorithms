#include <iostream>
#include "trees.h"

using namespace std;

bool is_with_in_bounds(Node* root, int upper, int lower) {
    if (root == nullptr) return true;

    if (lower > root->val || upper < root->val) return false;
    return is_with_in_bounds(root->left, root->val, lower) && is_with_in_bounds(root->right, upper, root->val);
}

bool validate(Node *root) {
    return is_with_in_bounds(root, INT_MAX, INT_MIN);
}

int main(int argc, char **argv) {
    Node* root = new Node(5);
    root->left = new Node(2);
    root->left->left = new Node(1);
    root->left->right = new Node(6);
    root->right = new Node(7);
    root->right->right = new Node(9);
    root->right->left = new Node(7);
    cout << validate(root) << "\n";
    return 0;
}
