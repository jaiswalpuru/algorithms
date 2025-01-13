#include <iostream>
#include <stack>

#include "trees.h"

using namespace std;

void inorder(Node* root, vector<int>& in) {
    if (root == nullptr) return;
    inorder(root->left, in);
    in.push_back(root->val);
    inorder(root->right, in);
}

int kth_small_recursive(Node* root, int k) {
    vector<int> in;
    inorder(root, in);
    return in[k-1];
}

int kth_small_iterative(Node* root, int k) {
    stack<Node*> st;
    Node* node = root;
    while (!st.empty() || node) {
        while (node) {
            st.push(node);
            node = node->left;
        }
        node = st.top();
        st.pop();
        k--;
        if (k == 0) return node->val;
        node = node->right;
    }
    return -1;
}

int main(int argc, char **argv) {
    Node* root = new Node(5);
    root->left = new Node(2);
    root->left->left = new Node(1);
    root->left->right = new Node(4);
    root->right = new Node(7);
    root->right->right = new Node(9);
    root->right->left = new Node(6);
    cout << kth_small_iterative(root, 5) << "\n";
    return 0;
}
