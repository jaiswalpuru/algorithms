#include <iostream>
#include <vector>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

typedef std::vector<int> vi;
typedef long long ll;

int recur(TreeNode *root, vi& res) {
    if (root == nullptr) return 0;
    int left = recur(root->left, res);
    int right = recur(root->right, res);
    res.push_back(root->val+left+right);
    return root->val+left+right;
}

int max_product_of_splitted_binary_tree(TreeNode* root) {
    vi res;
    int sum = recur(root, res);
    ll ans = 0;
    for (int i=0; i<res.size(); i++) {
        ans = std::max(ans, (ll)res[i]*(sum-res[i]));
    }
    return ans%((ll)1e9+7);
}

int main() {
    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->left = new TreeNode(4);
    root->left->right = new TreeNode(5);
    root->right->left = new TreeNode(6);
    std::cout<<max_product_of_splitted_binary_tree(root)<<"\n";
}
