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
    bool isEvenOddTree(TreeNode* root) {
        bool level = true; //true even level 
        queue<TreeNode*> q;
        q.push(root);

        while (!q.empty()) {
            int n = q.size();
            queue<TreeNode*> temp;
            for (int i = 0; i < n-1; i++) {
                TreeNode *curr = q.front();
                q.pop();
                TreeNode *next = q.front();
                if (level) {
                    if (curr->val >= next->val || curr->val % 2 == 0 || next->val % 2 == 0) return false;
                } else {
                    if (curr->val <= next->val || curr->val % 2 != 0 || next->val % 2 != 0) return false;
                }
                if (curr->left != nullptr) temp.push(curr->left);
                if (curr->right != nullptr) temp.push(curr->right);
            }
            if (level && !q.empty() && q.front()->val % 2 == 0) return false;
            if (!level && !q.empty() && q.front()->val % 2 != 0) return false;
            if (!q.empty() && q.front()->left != nullptr) temp.push(q.front()->left);
            if (!q.empty() && q.front()->right != nullptr) temp.push(q.front()->right);
            level = !level;
            q = temp;
        }
        return true;
    }
};
