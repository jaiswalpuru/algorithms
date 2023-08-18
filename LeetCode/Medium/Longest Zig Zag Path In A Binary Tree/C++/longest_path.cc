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
    int max_len;
    int longestZigZag(TreeNode* root) {
        max_len = 0;
        max_zig_zag(root, -1, 0);
        return max_len;
    }

    void max_zig_zag(TreeNode* root, int dir, int curr_len) {
        if (root == nullptr) return;
        max_len = max(max_len, curr_len);
        if (dir == 0) {
            max_zig_zag(root->right, 1, curr_len+1);
            max_zig_zag(root->left, 0, 1);
        } else {
            max_zig_zag(root->left, 0, curr_len+1);
            max_zig_zag(root->right, 1, 1);
        }
    }
};
