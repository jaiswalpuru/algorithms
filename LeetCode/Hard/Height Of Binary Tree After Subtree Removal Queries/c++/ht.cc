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
    int max_height_after_removal[(int)1e5 + 1];
    int curr_max_ht = 0;

    vector<int> treeQueries(TreeNode* root, vector<int>& queries) {
        vector<int> res;
        traverse_left_to_right(root, 0);
        curr_max_ht = 0;
        traverse_right_to_left(root, 0);
        
        for (int i = 0; i < queries.size(); i++) {
            res.push_back(max_height_after_removal[queries[i]]);
        }

        return res;
    }

    void traverse_right_to_left(TreeNode* root, int curr_ht) {
        if (root == nullptr) return;
        max_height_after_removal[root->val] = max(max_height_after_removal[root->val], curr_max_ht);
        curr_max_ht = max(curr_max_ht, curr_ht);
        traverse_right_to_left(root->right, curr_ht + 1);
        traverse_right_to_left(root->left, curr_ht + 1);
    }

    void traverse_left_to_right(TreeNode* root, int curr_ht) {
        if (root == nullptr) return;
        max_height_after_removal[root->val] = curr_max_ht;
        curr_max_ht = max(curr_max_ht, curr_ht);
        traverse_left_to_right(root->left, curr_ht + 1);
        traverse_left_to_right(root->right, curr_ht + 1);
    }

    int height(TreeNode* root, int val) {
        if (root == nullptr) return 0;
        if (root->val == val) return 0;
        int l = height(root->left, val);
        int r = height(root->right, val);
        return 1 + max(l, r);
    }
};
