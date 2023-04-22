#include <iostream>
#include <vector>

#define PB push_back

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

typedef std::vector<int> vi;

void traverse(TreeNode *r, vi &v) {
    if (r == NULL) return;
    if (r->left == NULL && r->right == NULL) v.PB(r->val);
    traverse(r->left, v);
    traverse(r->right, v);
}

bool leaf_similar(TreeNode* r1, TreeNode* r2) {
    vi v1, v2;
    traverse(r1, v1);
    traverse(r2, v2);
    if (v1.size() != v2.size()) return false;
    for (int i=0; i<v1.size(); i++) if (v1[i] != v2[i]) return false;
    return true;
}

int main () {
    TreeNode *r1 = new TreeNode(3);
    r1->left = new TreeNode(1);
    r1->right = new TreeNode(5);
    TreeNode *r2 = new TreeNode(2);
    r2->left = new TreeNode(1);
    r2->right = new TreeNode(5);
    std::cout<<leaf_similar(r1, r2)<<"\n";
}
