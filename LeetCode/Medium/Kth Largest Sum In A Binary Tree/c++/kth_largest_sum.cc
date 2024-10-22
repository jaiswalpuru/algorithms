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
    long long kthLargestLevelSum(TreeNode* root, int k) {
        queue<TreeNode*> q;
        vector<long long> sum_data;
        q.push(root);
        long long sum = 0LL;

        while (!q.empty()) {
            long long size = q.size();
            sum = 0LL;
            for (int i = 0; i < size; i++) {
                TreeNode* tmp = q.front();
                q.pop();
                sum += (long long)tmp->val;
                if (tmp->left != nullptr) q.push(tmp->left);
                if (tmp->right != nullptr) q.push(tmp->right);
            }
            sum_data.push_back(sum);
        }

        if (sum_data.size() < k) return -1;
        sort(sum_data.begin(), sum_data.end(), std::greater<long long>());

        return sum_data[k-1];
    }
};
