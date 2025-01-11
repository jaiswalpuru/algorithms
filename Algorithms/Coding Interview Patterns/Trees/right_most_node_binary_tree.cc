#include <iostream>
#include <queue>
#include <vector>
#include "trees.h"

using namespace std;

vector<int> right_most_node(Node* root) {
    vector<int> res;
    if (root == nullptr) return res;
    
    queue<Node*> q;
    q.push(root);

    while (!q.empty()) {
        int size = q.size();
        for (int i = 0; i < size; i++) {
            Node* curr = q.front();
            q.pop();
            if (i == size - 1) res.push_back(curr->val);
            if (curr->left) q.push(curr->left);
            if (curr->right) q.push(curr->right);
        }
    }

    return res;
}


int main(int argc, char **argv) {
    Node* root = new Node;
    root->val = 1;
    root->left = new Node;
    root->left->val = 2;
    root->left->left = new Node;
    root->left->left->val = 4;
    root->left->left->left = new Node;
    root->left->left->left->val = 8;
    root->left->right = new Node;
    root->left->right->val = 5;
    root->left->right->right = new Node;
    root->left->right->right->val = 11;
    root->left->left->right = new Node;
    root->left->left->right->val = 9;
    root->right = new Node;
    root->right->val = 3;
    root->right->left = new Node;
    root->right->left->val = 6;
    vector<int> res = right_most_node(root);
    for (int val : res) cout << val << " ";
    cout << "\n";
    return 0;
}
