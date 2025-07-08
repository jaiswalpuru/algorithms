/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string s = "";
        preorder(root, s);
        return s;
    }

    void preorder(TreeNode* root, string& s) {
        if (!root) {
            s += "null,";
            return;
        }
        
        s += to_string(root->val) + ",";
        preorder(root->left, s);
        preorder(root->right, s);
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        vector<string> nodes;
        stringstream ss(data);
        string t;
        while (getline(ss, t, ',')) nodes.push_back(t);
        int i = 0;
        return make_tree(nodes, i);
    }

    TreeNode* make_tree(vector<string>& nodes, int& i) {
        if (i == nodes.size()) return nullptr;
        if (nodes[i] == "null") {
            i++;
            return nullptr;
        }
        TreeNode* root = new TreeNode(stoi(nodes[i]));
        i++;
        root->left = make_tree(nodes, i);
        root->right = make_tree(nodes, i);
        return root;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec ser, deser;
// TreeNode* ans = deser.deserialize(ser.serialize(root));
