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
    TreeNode* replaceValueInTree(TreeNode* root) {
        if (root == nullptr) return root;
        queue<TreeNode*> q;
        q.push(root);
        int prev_lvl_sum = root->val;

        while (!q.empty()) {
            int size = q.size();
            int curr_lvl_sum = 0;

            for (int i = 0; i < size; i++) {
                TreeNode* curr = q.front();
                q.pop();

                curr->val = prev_lvl_sum - curr->val;
                int sibling_sum = (curr->left != nullptr ? curr->left->val : 0) + (curr->right != nullptr ? curr->right->val : 0);
                if (curr->left != nullptr) {
                    curr_lvl_sum += curr->left->val;
                    curr->left->val = sibling_sum;
                    q.push(curr->left);
                }
                if (curr->right != nullptr) {
                    curr_lvl_sum += curr->right->val;
                    curr->right->val = sibling_sum;
                    q.push(curr->right);
                }
            }
            prev_lvl_sum = curr_lvl_sum;
        }

        return root;
    }
};
