/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* increasingBST(TreeNode* root) {
        vector<TreeNode*> in;
        inorder(root, in);
        for (int i=1; i<in.size(); i++) {
            in[i-1]->left = nullptr;
            in[i-1]->right = in[i];
        }
        in[in.size()-1]->left = nullptr;
        in[in.size()-1]->right = nullptr;
        return in[0];
    }

    void inorder(TreeNode* root, vector<TreeNode*>& in) {
        if (root == nullptr) return;
        inorder(root->left, in);
        in.push_back(root);
        inorder(root->right, in);
    }
};
