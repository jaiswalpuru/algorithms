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
    int goodNodes(TreeNode* root) {
        int good_nodes = 0;
        dfs(root, root, good_nodes);
        return good_nodes;
    }

    void dfs(TreeNode* root, TreeNode* curr, int& good_nodes) {
        if (!root) return;
        if (root->val >= curr->val) {
            good_nodes++;
            dfs(root->left, root, good_nodes);
            dfs(root->right, root, good_nodes);
        } else {
            dfs(root->left, curr, good_nodes);
            dfs(root->right, curr, good_nodes);
        }
    }
};
