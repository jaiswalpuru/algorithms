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
    unordered_map<int, int> hash_map;
    int pre_ind;
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        for (int i = 0; i < inorder.size(); i++) hash_map[inorder[i]] = i;
        pre_ind = 0;
        return _build_tree(preorder, 0, preorder.size() - 1);
    }

    TreeNode* _build_tree(vector<int>& preorder, int l, int r) {
        if (l > r) return nullptr;
        int root_val = preorder[pre_ind++];
        TreeNode* root = new TreeNode(root_val);
        root->left = _build_tree(preorder, l, hash_map[root_val] - 1);
        root->right = _build_tree(preorder, hash_map[root_val] + 1, r);
        return root;
    }
};
