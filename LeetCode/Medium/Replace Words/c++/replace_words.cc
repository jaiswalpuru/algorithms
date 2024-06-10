struct Node {
    string word;
    Node* children[26];
};

class Trie {
    Node* t;

public:
    Trie() {
        t = new Node();
    }

    string search(string word) {
        Node* root = this->t;
        int ind;

        for (int i = 0; i < word.length(); i++) {
            ind = word[i] - 'a';
            if (root->children[ind] == nullptr) break;
            if (root->word != "") return root->word;
            root = root->children[ind];
        }

        return root->word;

    }

    void insert(string word) {
        Node* root = this->t;
        int ind;

        for (int i = 0; i < word.length(); i++) {
            ind = word[i] - 'a';
            if (root->children[ind] == nullptr) root->children[ind] = new Node();
            root = root->children[ind];
        }
        root->word = word;
    }

};

class Solution {
public:
    string replaceWords(vector<string>& dictionary, string sentence) {
        Trie *trie = new Trie();
        stringstream ss(sentence);
        vector<string> words;
        string s, temp;

        for (string s : dictionary) trie->insert(s);
        while (getline(ss, s, ' ')) words.push_back(s);

        s = "";
        for (int i = 0; i < words.size(); i++) {
            temp = trie->search(words[i]);
            if (temp == "") s += words[i];
            else s += temp;
            s += " ";
        }

        return s.substr(0, s.length()-1);
    }
};

