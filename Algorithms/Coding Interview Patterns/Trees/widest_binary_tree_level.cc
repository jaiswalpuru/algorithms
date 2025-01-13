#include <iostream>
#include <queue>
#include "trees.h"

using namespace std;

int widest_level(Node* root) {
    if (root == nullptr) return 0;
    
    queue<pair<Node*, int>> q;
    q.push({root, 0});

    int widest_lvl = 0;

    while (!q.empty()) {
        int size = q.size();
        int left_most = q.front().second;
        int right_most = left_most;

        for (int i = 0; i < size; i++){
            auto curr = q.front();
            q.pop();
            if (curr.first->left) q.push({curr.first->left, 2 * curr.second + 1});
            if (curr.first->right) q.push({curr.first->right, 2 * curr.second + 2});
            right_most = curr.second;
        }

        widest_lvl = max(widest_lvl, right_most - left_most + 1);
    }
    return widest_lvl;
}

int main(int argc, char **argv) {
    Node *root = new Node(1);
    root->left = new Node(2);
    root->left->left = new Node(4);
    root->left->left->left = new Node(8);
    root->left->left->right = new Node(9);
    root->left->right = new Node(5);
    root->left->right->right = new Node(11);
    root->right = new Node(3);
    root->right->right = new Node(7);
    root->right->right->left = new Node(14);
    cout << widest_level(root) << "\n";
    return 0;
}
