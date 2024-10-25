class Solution {
    struct Node {
        bool is_leaf;
        unordered_map<string, Node*> children;
        Node() : is_leaf(false) {}
    };

    void delete_trie(Node* root) {
        if (root == nullptr) return;
        for (auto& pair : root->children) {
            delete_trie(pair.second);
        }
        delete root;
    }

public:
    Node* root;
    Solution() : root(new Node()) {}
    ~Solution() { delete_trie(root); }

    void insert_into_trie(vector<string>& folder) {
        for (string& path : folder) {
            Node* curr = root;
            istringstream iss(path);
            string folder_name;
            while (getline(iss, folder_name, '/')) {
                if (folder_name.empty()) continue;
                if (curr->children.find(folder_name) == curr->children.end())curr->children[folder_name] = new Node();
                curr = curr->children[folder_name];
            }
            curr->is_leaf = true;
        }
    }

    vector<string> removeSubfolders(vector<string>& folder) {
        vector<string> res;
        insert_into_trie(folder);
        
        for (string& path : folder) {
            Node* curr = root;
            istringstream iss(path);
            string folder_name;
            bool is_sub = false;

            while (getline(iss, folder_name, '/')) {
                if (folder_name.empty()) continue;
                Node* next = curr->children[folder_name];
                if (next->is_leaf && iss.rdbuf()->in_avail() != 0) {
                    is_sub = true;
                    break;
                }
                curr = next;
            }
            if (!is_sub) res.push_back(path);
        }

        return res;  
    }
};
