/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
typedef struct TreeNode TreeNode;

bool is_same_tree(TreeNode* p, TreeNode* q) {
    if (p == NULL && q == NULL) return true;
    if (p == NULL || q == NULL) return false;
    if (p->val != q->val) return false;
    return is_same_tree(p->left, q->left) && is_same_tree(p->right, q->right);
}

bool isSubtree(struct TreeNode* root, struct TreeNode* subRoot){
    if (root == NULL) return false;
    if (root->val == subRoot->val) if (is_same_tree(root, subRoot)) return true;
    return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);
}
