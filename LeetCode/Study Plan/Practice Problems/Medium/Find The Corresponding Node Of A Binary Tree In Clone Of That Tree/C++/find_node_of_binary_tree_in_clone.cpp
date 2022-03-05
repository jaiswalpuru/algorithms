#include <iostream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

TreeNode *get_node(TreeNode *cloned, TreeNode *target, TreeNode *refer) {
    if (cloned == NULL) {
        return NULL;
    }

    if (cloned->val == target->val) {
        refer = cloned;
        return refer;
    }

    TreeNode *left = get_node(cloned->left, target, refer);
    if (left) {
        return left;
    }
    return get_node(cloned->right, target, refer);
}


TreeNode* get_copy(TreeNode *original, TreeNode *clone, TreeNode *target) {
    TreeNode *refer;
    return get_node(clone, target, refer);
}


int main() {
    TreeNode *tree = new TreeNode(4);
    tree->left = new TreeNode(4);
    tree->right = new TreeNode(3);
    tree->right->right = new TreeNode(19);
    tree->right->left = new TreeNode(6);
    
    TreeNode *cloned = new TreeNode(7);
    cloned->left = new TreeNode(4);
    cloned->right = new TreeNode(3);
    cloned->right->right = new TreeNode(19);
    cloned->right->left = new TreeNode(6);

    TreeNode *r;
    r = get_copy(tree, cloned, tree->right);
    cout<<r->val<<endl;
    return 0;
}