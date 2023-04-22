#include <iostream>
#include <vector>

typedef std::vector<char> vc;
typedef std::vector<vc> vcc;
typedef std::vector<std::string> vs;
typedef std::vector<bool> vb;
typedef std::vector<vb> vbb;

//--------------------------efficient approach--------------------------

typedef struct Trie {
    Trie *children[26];
    std::string word;
} T;

T *new_node() {
    T *node = new T;
    node->word = "";
    for (int i=0; i<26; i++) {
        node->children[i] = NULL;
    }
    return node;
}

void insert(T *node, std::string word) {
    T *temp = node;
    for (int i=0;i<word.length(); i++) {
        int index = word[i] - 'a';
        if (!temp->children[index]) temp->children[index] = new_node();
        temp = temp->children[index];
    }
    temp->word = word;
}

void _dfs(vcc &board, T *node, int i, int j, vbb &visited, vs &res) {
    if (i<0 || j<0 || i>=board.size() || j>=board[0].size() || visited[i][j]) return;
    int ch = board[i][j]-'a';
    if (node->children[ch] == NULL) return;
    node = node->children[ch];
    if (node->word.length() > 0) {
        res.push_back(node->word);
        node->word = "";
    }
    visited[i][j] = true;
    _dfs(board, node, i+1, j, visited, res);
    _dfs(board, node, i-1, j, visited, res);
    _dfs(board, node, i, j+1, visited, res);
    _dfs(board, node, i, j-1, visited, res);
    visited[i][j] = false;
}

vs word_search_two_eff(vcc &board, vs &words) {
    vs res;
    int n=board.size(), m = board[0].size();
    T *trie = new_node();
    vbb visited(n, vb(m));
    for (int i=0;i<words.size(); i++) {
        insert(trie, words[i]);
    }
    for (int i=0;i<n;i++) {
        for (int j=0;j<m;j++) {
            _dfs(board, trie, i, j, visited, res);
        }
    }
    return res;
}
//--------------------------efficient approach--------------------------

//--------------------brute force----------------(try each possibility)
bool dfs(vcc board, int i, int j, int k, std::string &word, vbb &visited);

vs word_search_two(vcc &board, vs &words) {
    vs res;
    int n = board.size(), m = board[0].size();
    vbb visited(n, vb(m));
    for (int k=0;k<words.size(); k++) {
        bool f = false;
        for (int i=0;i<n;i++) {
            for (int j=0;j<m;j++) {
                if (board[i][j] == words[k][0]) {
                    if (dfs(board, i,j, 0, words[k], visited)){
                        res.push_back(words[k]);
                        f = true;
                        break;
                    }
                }
            }
            if (f) break;
        }
    }
    return res;
}

bool dfs(vcc board, int i, int j, int k, std::string &word, vbb &visited) {
    if (k == word.length()) return true;
    if (i<0 || j<0 || i>=board.size() || j>=board[0].size() || visited[i][j])
        return false;
    if (board[i][j] != word[k]) return false;
    visited[i][j] = true;
    bool res =
        dfs(board, i+1, j, k+1, word, visited) || 
        dfs(board, i-1, j, k+1, word, visited) ||
        dfs(board, i, j+1, k+1, word, visited) ||
        dfs(board, i, j-1, k+1, word, visited);
    visited[i][j] = false;
    return res;
}
//--------------------brute force----------------(try each possibility)

int main() {
    vcc board{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}};
    vs words{"oath","pea","eat","rain"};
    vs res = word_search_two_eff(board, words);
    for (auto it=res.begin(); it!=res.end(); it++) std::cout<<*it<<" ";
    std::cout<<"\n";
}
