#include <iostream>
#include <queue>
#include <stack>
#include "trees.h"

using namespace std;

void bfs(Node* root);
void invert_recursive(Node** root);
void invert_iterative(Node** root);

int main(int argc, char **argv) {
    Node* root = new Node;
    root->val = 5;
    root->left = new Node;
    root->right = new Node;
    root->left->val = 1;
    root->right->val = 8;
    root->left->left = new Node;
    root->left->left->val = 7;
    root->left->right = new Node;
    root->left->right->val = 6;
    root->right->right = new Node;
    root->right->right->val = 4;
    cout << "before invert \n";
    bfs(root);
    invert_iterative(&root);
    cout << "after invert \n";
    bfs(root);
    return 0;
}

void bfs(Node* root) {
    if (root == nullptr) return;
    queue<Node*> q;
    q.push(root);
    
    while (!q.empty()) {
        int size = q.size();

        for (int i = 0; i < size; i++) {
            Node *curr = q.front();
            q.pop();
            cout << curr->val <<" ";
            if (curr->left) q.push(curr->left);
            if (curr->right) q.push(curr->right);
        }
        cout << "\n";
    }
}

void invert_recursive(Node** root) {
    if (*root == nullptr) return;
    Node* temp = (*root)->left;
    (*root)->left = (*root)->right;
    (*root)->right = temp;
    invert_recursive(&(*root)->left);
    invert_recursive(&(*root)->right);
}

void invert_iterative(Node** root) {
    if (*root == nullptr) return;
    Node* r = *root;
    stack<Node*> st;

    st.push(r);
    while (!st.empty()) {
        Node* curr = st.top();
        st.pop();
        Node* temp = curr->left;
        curr->left = curr->right;
        curr->right = temp;
        if (curr->left) st.push(curr->left);
        if (curr->right) st.push(curr->right);
    }
    *root = r;
}

