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
    bool isValidBST(TreeNode* root) {
        return is_valid(root, LLONG_MIN, LLONG_MAX);
    }

    bool is_valid(TreeNode* root, long long min, long long max) {
        if (!root) return true;
        if (min >= root->val || max <= root->val) return false;
        return is_valid(root->left, min, root->val) && is_valid(root->right, root->val, max);
    }
    
};
