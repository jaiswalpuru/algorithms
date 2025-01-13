#include <iostream>
#include <unordered_map>
#include <vector>

#include "trees.h"

using namespace std;

void inorder(Node* root) {
    if (root == nullptr) return;
    inorder(root->left);
    cout << root->val << " ";
    inorder(root->right);
}

Node* _build_tree(int left, int right, vector<int> in, vector<int> pre, int &pre_ind, unordered_map<int, int> index) {
    if (left > right) return nullptr;
    int v = pre[pre_ind];

    int in_v = index[v];
    Node* root = new Node(v);
    pre_ind++;
    root->left = _build_tree(left, in_v - 1, in, pre, pre_ind, index);
    root->right = _build_tree(in_v + 1, right, in, pre, pre_ind, index);
    return root;
}

Node* build_tree(vector<int> pre, vector<int> in) {
    unordered_map<int, int> index;
    int pre_ind = 0;
    for (int i = 0; i < in.size(); i++) index[in[i]] = i;
    return _build_tree(0, in.size() - 1, in, pre, pre_ind, index);
}

int main(int argc, char **argv) {
    Node *root = build_tree({5, 9, 2, 3, 4, 7}, {2, 9, 5, 4, 3, 7});
    inorder(root);
    cout << "\n";
    return 0;
}


