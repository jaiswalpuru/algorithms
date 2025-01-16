#include <iostream>
#include "tries.h"

using namespace std;

void dfs(vector<vector<char>> board, int r, int c, vector<string>& res, Node* node) {
    if (node->is_word) {
        res.push_back(node->word);
        node->is_word = false;
        node->word = "";
    }
    board[r][c] = '#';

    int dir[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    for (int i = 0; i < 4; i++) {
        int x = r + dir[i][0];
        int y = c + dir[i][1];
        if (x >= 0 && y >= 0 && x < board.size() && y < board[0].size()) {
            if (node->children.find(board[x][y]) != node->children.end()) {
                dfs(board, x, y, res, node->children[board[x][y]]);
            }
        }
    }
}

vector<string> find_all_words(vector<vector<char>> board, vector<string> words) {
    Trie* t = new Trie();
    vector<string> res;
    for (string s : words) t->insert(s);

    for (int i = 0; i < board.size(); i++) {
        for (int j = 0; j < board[0].size(); j++) {
            if (t->root->children.find(board[i][j]) != t->root->children.end()) dfs(board, i, j, res, t->root->children[board[i][j]]);
        }
    }
    return res;
}

int main(int argc, char** argv) {
    vector<vector<char>> board = {{'b', 'y', 's'}, {'r', 't', 'e'}, {'a', 'i', 'n'}};
    vector<string> words = {"byte", "bytes", "rat", "rain", "trait", "train"};

    vector<string> res = find_all_words(board, words);
    for (string s : res) cout << s << "\n";
    return 0;
}
