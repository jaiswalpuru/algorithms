/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

void recur(struct TreeNode* root, int val, int* good)
{
    if (root == NULL) return;
    if (root->left != NULL && val <= root->left->val)
    {
        (*good)++;
        recur(root->left, root->left->val, good);
    }
    else
    {
        recur(root->left, val, good);
    }

    if (root->right != NULL && val <= root->right->val)
    {
        (*good)++;
        recur(root->right, root->right->val, good);
    }
    else
    {
        recur(root->right, val, good);
    }
}

int goodNodes(struct TreeNode* root)
{
    int good = 1;
    recur(root, root->val, &good);
    return good;
}
