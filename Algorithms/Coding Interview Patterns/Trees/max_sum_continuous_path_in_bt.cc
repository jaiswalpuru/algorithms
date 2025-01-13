#include <iostream>
#include "trees.h"

using namespace std;

int res = INT_MIN;

int _max_path(Node* root) {
    if (root == nullptr) return 0;
    
    int left_sum = max(_max_path(root->left), 0);
    int right_sum = max(_max_path(root->right), 0);

    res = max(res, root->val + left_sum + right_sum);
    return root->val + max(left_sum, right_sum);
}

int max_path_sum(Node* root) {
    _max_path(root);
    return res;
}

int main(int argc, char **argv) {
    Node* root = new Node(5);
    root->left = new Node(-10);
    root->left->left = new Node(1);
    root->left->left->left = new Node(11);
    root->left->right = new Node(-7);
    root->left->right->left = new Node(-1);
    root->right = new Node(8);
    root->right->left = new Node(9);
    root->right->right = new Node(7);
    root->right->right->right = new Node(-3);
    root->right->right->left = new Node(6);
    cout << max_path_sum(root) << "\n";
    return 0;
}
