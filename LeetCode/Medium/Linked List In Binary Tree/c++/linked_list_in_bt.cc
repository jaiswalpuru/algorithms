/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
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
    bool isSubPath(ListNode* head, TreeNode* root) {
        return traverse(root, head);
    }

    bool traverse(TreeNode *root, ListNode *head) {
        if (root == nullptr) return false;
        if (head->val == root->val && check_path(root, head)) {
            return true;
        }
        return traverse(root->left, head) || traverse(root->right, head);
    }

    bool check_path(TreeNode *root, ListNode *head) {
        if (head == nullptr) return true;
        if (root == nullptr) return false;
        if (head->val != root->val) return false;
        return check_path(root->left, head->next) || check_path(root->right, head->next);
    }

};
