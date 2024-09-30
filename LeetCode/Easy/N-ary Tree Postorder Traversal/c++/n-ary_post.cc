/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val) {
        val = _val;
    }

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
public:
    vector<int> postorder(Node* root) {
        stack<Node*> st;
        vector<int> res;
        if (root == nullptr) return res;
        
        st.push(root);
        while (!st.empty()) {
            Node* curr = st.top();
            st.pop();
            res.push_back(curr->val);
            for (int i = 0; i < curr->children.size(); i++) st.push(curr->children[i]);
        }

        reverse(res.begin(), res.end());
        return res;

    }
};
