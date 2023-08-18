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
    int maxLevelSum(TreeNode* root) {
        queue<TreeNode*> q{{root}};
        int max_sum = INT_MIN;
        int lvl = 0;
        int res = 0;
        while(!q.empty()) {
            lvl++;
            int sum=0;
            int n = q.size();
            for (int i=0; i<n; i++) {
                TreeNode* curr = q.front();
                q.pop();
                sum += curr->val;
                if (curr->left != nullptr) q.push(curr->left);
                if (curr->right != nullptr) q.push(curr->right);
            }
            if (max_sum < sum) {
                max_sum = sum;
                res = lvl;
            }
        }
        return res;
    }
};
