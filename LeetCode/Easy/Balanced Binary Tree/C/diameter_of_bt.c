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

int height(TreeNode* root) {
    if (root == NULL) return 0;
    int left = height(root->left);
    int right = height(root->right);
    return 1+MAX(left, right);
}

bool isBalanced(struct TreeNode* root){
    if (root == NULL) return true;
    int left = height(root->left);
    int right = height(root->right);
    if (abs(left-right) > 1) return false;
    return isBalanced(root->left) && isBalanced(root->right);
}
