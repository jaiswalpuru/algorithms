#include <iostream>
#include <string>
#include <vector>
#include <numeric>
#include <sstream>
#include "trees.h"

using namespace std;

Node* build_tree(vector<string> s, int& i) {
    string val = s[i];
    i++;
    if (val == "#") return nullptr;

    Node* root = new Node(stoi(val));
    root->left = build_tree(s, i);
    root->right = build_tree(s, i);
    return root;
}

Node* deserialize(string s) {
    vector<string> st;
    stringstream ss(s);
    string token;
    string delimiter = ",";
    int i = 0;

    while(getline(ss, token, delimiter.front())) st.push_back(token);
    return build_tree(st, i);
}

void preorder(Node* root, vector<string>& s) {
    if (root == nullptr) {
        s.push_back("#");
        return;
    }
    s.push_back(to_string(root->val));
    preorder(root->left, s);
    preorder(root->right, s);
}

string serialize(Node *root) {
    vector<string> s;
    preorder(root, s);
    string res = "";
    for (int i = 0; i < s.size(); i++) {
        res += s[i];
        if (i != s.size() - 1) res += ",";
    }
    return res;
}

void inorder(Node* root) {
    if (root == nullptr) return;
    inorder(root->left);
    cout << root->val << " ";
    inorder(root->right);
}

int main(int argc, char **argv) {
    Node* root = new Node(5);
    root->left = new Node(9);
    root->right = new Node(3);
    inorder(root);
    cout << "\n";
    string s = serialize(root);
    cout << "serialized string : " << s <<"\n";
    Node* node = deserialize(s);
    cout << "deserializing ... \n"; 
    inorder(node);
    cout << "\n";
    return 0;
}
