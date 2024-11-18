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
    int l = INT_MAX;
    int r = INT_MIN;
    vector<vector<int>> verticalTraversal(TreeNode* root) {
        unordered_map<int, vector<pair<int, int>>> u_map;
        vertical_order_traversal(root, u_map, 0, 0);
        vector<vector<int>> res;

        for (int i = l; i <= r; i++) {
            vector<pair<int, int>> t = u_map[i];
            sort(t.begin(), t.end(), [&](const auto& a, const auto& b) {
                if (a.second == b.second) {
                    return a.first < b.first;
                }
                return a.second < b.second;
            });
            vector <int> temp;
            for (int j = 0; j < t.size(); j++) temp.push_back(t[j].first);
            res.push_back(temp);
        }

        return res;
    }

    void vertical_order_traversal(TreeNode* root, unordered_map<int, vector<pair<int, int>>>& u_map, int i, int k) {
        if (root == nullptr) return;
        l = min(l, i);
        r = max(r, i);
        u_map[i].push_back({root->val, k});
        vertical_order_traversal(root->left, u_map, i - 1, k + 1);
        vertical_order_traversal(root->right, u_map, i + 1, k + 1);
    }
};
