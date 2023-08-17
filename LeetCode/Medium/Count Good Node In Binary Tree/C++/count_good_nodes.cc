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
        int cnt=1;
        recur(root, root->val, cnt);
        return cnt;
    }

    void recur(TreeNode* root, int val, int& cnt) {
        if (root == nullptr) return;
        if (root->left != nullptr && root->left->val >= val) {
            cnt++;
            recur(root->left, root->left->val, cnt);
        } else recur(root->left, val, cnt);

        if (root->right != nullptr && root->right->val >= val) {
            cnt++;
            recur(root->right, root->right->val, cnt);
        } else recur(root->right, val, cnt);
    }
};
