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
    bool isCousins(TreeNode* root, int x, int y) {
        queue<pair<int, TreeNode*>> q;
        q.push({-1, root});

        while (!q.empty()) {
            int size = q.size();
            int parent_x = 0;
            int parent_y = 0;
            bool found_x = false;
            bool found_y = false;
            for (int i = 0; i < size; i++) {
                int parent_val = q.front().first;
                TreeNode* curr = q.front().second;
                q.pop();

                if (x == curr->val || y == curr->val && parent_val == -1) return false;
                
                if (curr->left != nullptr) {
                    q.push({curr->val, curr->left});
                    if (curr->left->val == x) {
                        parent_x = curr->val;
                        found_x = true;
                    }
                    if (curr->left->val == y) {
                        parent_y = curr->val;
                        found_y = true;
                    }
                }
                
                if (curr->right != nullptr) {
                    q.push({curr->val, curr->right});
                    if (curr->right->val == x) {
                        parent_x = curr->val;
                        found_x = true;
                    }
                    if (curr->right->val == y) {
                        parent_y = curr->val;
                        found_y = true;
                    }
                }
                
                if (found_x && found_y && parent_x == parent_y) return false;
                if (found_x && found_y) return true;
            }
        }
        return false;
    }
};
