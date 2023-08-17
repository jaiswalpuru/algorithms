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
    int pathSum(TreeNode* root, int targetSum) {
        map<long, int> sum_map;
        int cnt = 0;
        recur(root, targetSum, 0, sum_map, cnt);
        return cnt;
    }

    void recur(TreeNode* root, int targetSum, long sum, map<long,int>& sum_map, int& cnt) {
        if (root == nullptr) return;
        sum += root->val;
        if (sum == targetSum) cnt++;
        cnt += sum_map[sum-targetSum];
        sum_map[sum]++;
        recur(root->left, targetSum, sum, sum_map, cnt);
        recur(root->right, targetSum, sum, sum_map, cnt);
        sum_map[sum]--;
    }
};
