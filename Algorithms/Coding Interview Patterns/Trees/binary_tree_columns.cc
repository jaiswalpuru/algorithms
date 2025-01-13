#include <iostream>
#include <unordered_map>
#include <queue>
#include "trees.h"

using namespace std;

vector<vector<int>> column(Node* root) {
    vector<vector<int>> res;
    if (root == nullptr) return res;
    queue<pair<Node*, int>> q;
    unordered_map<int, vector<int>> mp;

    q.push({root, 0});
    while (!q.empty()) {
        int size = q.size();
        for (int i = 0; i < size; i++) {
            auto curr = q.front();
            q.pop();
            mp[curr.second].push_back(curr.first->val);
            if (curr.first->left) q.push({curr.first->left, curr.second - 1});
            if (curr.first->right) q.push({curr.first->right, curr.second + 1});
        }
    }

    for (auto val : mp) {
        vector<int> t;
        for (int v : val.second) t.push_back(v);
        res.push_back(t);
    }
    return res;
}


int main(int argc, char **argv) {
    Node* root = new Node(5);
    root->left = new Node(9);
    root->right = new Node(3);
    root->left->left = new Node(2);
    root->left->right = new Node(1);
    root->right->left = new Node(4);
    root->right->right = new Node(7);
    vector<vector<int>> res = column(root);
    for (int i = 0; i < res.size(); i++) {
        for (int j = 0; j < res[i].size(); j++) {
            cout << res[i][j] << " ";
        }
        cout << "\n";
    }
    return 0;
}
