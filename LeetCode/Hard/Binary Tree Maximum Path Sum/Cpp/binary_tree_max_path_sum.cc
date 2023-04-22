#include <iostream>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode():val(0), left(nullptr), right(nullptr) {}
    TreeNode(int val):val(val), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {} 
};

int recur(TreeNode *root, int &res) {
    if (root == nullptr) return 0;
    int left = recur(root->left, res);
    int right = recur(root->right, res);
    int l_sum = std::max(root->val, root->val+left);
    int r_sum = std::max(root->val, root->val+right);
    int total_sum = l_sum + r_sum - root->val;
    res = std::max(res, total_sum);
    return std::max(l_sum, r_sum);
}

int binary_tree_max_path_sum(TreeNode *root) {
    if (root == nullptr) return 0;
    int res = INT_MIN;
    recur(root, res);
    return res;
}

int main() {
    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    std::cout<<binary_tree_max_path_sum(root)<<"\n";
}
