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
    TreeNode* balanceBST(TreeNode* root) {
        vector<int> tree_nodes;
        inorder(root, tree_nodes);
        return construct(tree_nodes, 0, tree_nodes.size()-1);
    }

    void inorder(TreeNode* root, vector<int> &tree_nodes) {
        if (root == nullptr) return;

        inorder(root->left, tree_nodes);
        tree_nodes.push_back(root->val);
        inorder(root->right, tree_nodes);
    }

    TreeNode* construct(vector<int> tree_nodes, int st, int end) {
        if (st > end) return nullptr;

        int mid = st + (end-st)/2;
        TreeNode *left_subtree = construct(tree_nodes, st, mid-1);
        TreeNode *right_subtree = construct(tree_nodes, mid+1, end);
        return new TreeNode(tree_nodes[mid], left_subtree, right_subtree);
    }
};
