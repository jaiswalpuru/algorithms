#include <iostream>
#include <queue>
#include "trees.h"

using namespace std;

void bfs(Node* root) {
    if (root == nullptr) return;

    queue<Node*> q;
    q.push(root);
    
    while (!q.empty()) {
        Node *curr = q.front();
        q.pop();
        cout << curr->val << " ";
        if (curr->left) q.push(curr->left);
        if (curr->right) q.push(curr->right);
    }
}


int main(int argc, char **argv) {
    Node* root = new Node;
    root->val = 1;
    root->left = new Node;
    root->left->val = 2;
    root->right = new Node;
    root->right->val = 3;
    bfs(root);
    cout << "\n";
    return 0;
}
