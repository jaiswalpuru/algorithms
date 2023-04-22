#include <iostream>
#include <queue>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

typedef std::queue<TreeNode*> qt;

int count_nodes(TreeNode *root) {
    if (root == nullptr) return 0;
    qt q;
    int cnt = 0;
    q.push(root);
    while(!q.empty()) {
        auto curr = q.front();
        q.pop();
        cnt++;
        if (curr->left) q.push(curr->left);
        if (curr->right) q.push(curr->right);
    }
    return cnt;
}

int main() {
    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->left = new TreeNode(4);
    root->left->right = new TreeNode(5);
    root->right->left = new TreeNode(6);
    std::cout<<count_nodes(root)<<"\n";
}
