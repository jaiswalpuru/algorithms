/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
#define MAX(a, b) (((a) > (b)) ? (a) : (b))

typedef struct TreeNode TreeNode;

int recur(TreeNode* root, int *res) {
    if (root == NULL) return 0;
    int left = recur(root->left, res);
    int right = recur(root->right, res);
    int l_sum = MAX(root->val, left + root->val);
    int r_sum = MAX(root->val, right + root->val);
    int total = l_sum + r_sum - root->val;
    *res = MAX(*res, total);
    return MAX(l_sum, r_sum);
}

int maxPathSum(TreeNode* root){
    int res = -(1e9+7);
    recur(root, &res);
    return res;
}
