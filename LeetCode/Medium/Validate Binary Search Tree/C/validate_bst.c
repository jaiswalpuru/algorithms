/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
typedef struct TreeNode TreeNode;

bool recur(TreeNode* root, double lower_bound, double upper_bound) {
    return (root == NULL) ||
        (
            ((double)root->val > lower_bound && (double)root->val < upper_bound) &&
            recur(root->left, lower_bound, (double)root->val) &&
            recur(root->right, (double)root->val, upper_bound)
        );
}

bool isValidBST(TreeNode* root){
    return recur(root, -INFINITY, INFINITY);
}
