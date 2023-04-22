/**
 * Definition for a binary tree node.
 * struct Node {
 *     char val;
 *     Node *left;
 *     Node *right;
 *     Node() : val(' '), left(nullptr), right(nullptr) {}
 *     Node(char x) : val(x), left(nullptr), right(nullptr) {}
 *     Node(char x, Node *left, Node *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    bool checkEquivalence(Node* root1, Node* root2) {
        unordered_map <char, int> m1,m2;
        inorder(root1, m1);
        inorder(root2, m2);
        for (auto x : m1) {
            if (m2[x.first] != x.second) {
                return false;
            }
        }
        for (auto x : m2) {
            if (m1[x.first] != x.second) {
                return false;
            }
        }
        return true;
    }
    
    void inorder(Node* r, unordered_map<char, int> &m) {
        if (r==NULL) {
            return;
        }
        inorder(r->left, m);
        m[r->val]++;
        inorder(r->right, m);
    }
};