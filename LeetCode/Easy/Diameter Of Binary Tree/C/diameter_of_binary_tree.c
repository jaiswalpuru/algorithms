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

int recur(TreeNode* root, int* d) {
    if (root == NULL) return 0;
    int left = recur(root->left, d);
    int right = recur(root->right, d);
    *d = MAX(*d, left+right);
    return 1 + MAX(left, right);
}

int diameterOfBinaryTree(TreeNode* root) {
    int d = 0;
    recur(root, &d);
    return d;
}

