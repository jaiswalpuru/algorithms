#include <iostream>
#include "trees.h"

using namespace std;

Node* lca(Node *root, Node* p, Node* q) {
    if (root == nullptr) return root;
    
    if (root == p || root == q) return root;
    Node* left = lca(root->left, p, q);
    Node* right = lca(root->right, p, q);
    if (left == nullptr) return right;
    if (right == nullptr) return left;
    return root;
}

int main(int argc, char **argv) {
    Node* root = new Node(1);
    root->left = new Node(2);
    root->left->left = new Node(4);
    root->left->right = new Node(5);
    root->right = new Node(3);
    root->right->right = new Node(7);
    root->right->left = new Node(6);
    root->right->left->left = new Node(8);
    root->right->left->right = new Node(9);
    cout << lca(root, root->right->left->left, root->right->right)->val << "\n";
    return 0;
}
