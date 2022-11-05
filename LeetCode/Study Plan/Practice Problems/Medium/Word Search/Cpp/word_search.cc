#include <iostream>
#include <vector>

typedef std::vector<char> vc;
typedef std::vector<vc> vcc;
typedef std::vector<bool> vb;
typedef std::vector<vb> vbb;

bool dfs(std::string word, int i, int j, int k, vbb &visited, vcc &board) {
    if (k == word.length()) return true;
    if (i<0 || j<0 || i>=board.size() || j>=board[0].size() || visited[i][j]) return false;
    if (word[k] != board[i][j]) return false;
    visited[i][j] = true;
    bool res = 
        dfs(word, i+1, j, k+1, visited, board) ||
        dfs(word, i-1, j, k+1, visited, board) ||
        dfs(word, i, j+1, k+1, visited, board) ||
        dfs(word, i, j-1, k+1, visited, board);
    visited[i][j] = false;
    return res;
}

bool word_search(vcc &board, std::string word) {
    int n = board.size(), m = board[0].size();
    vbb visited(n, vb(m));
    for (int i=0;i<n;i++) {
        for (int j=0;j<m; j++) {
            if (board[i][j] == word[0]) {
                if (dfs(word, i, j, 0, visited, board)) {
                    return true;
                }
            }
        }
    }
    return false;
}

int main () {
    vcc v{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}};
    std::cout<<word_search(v, "ABCCED")<<"\n";
}
