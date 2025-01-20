#include <iostream>
#include <vector>

using namespace std;

vector<vector<int>> dir = {{0, -1}, {0, 1}, {1, 0}, {-1, 0}};

int dfs(int r, int c, vector<vector<int>> board, vector<vector<int>> memo) {
    if (memo[r][c] != 0) return memo[r][c];
    int path = 1;
    for (int i = 0; i < 4; i++) {
        int x = r + dir[i][0];
        int y = c + dir[i][1];
        if (x >=0 && y >= 0 && x < board.size() && y < board[0].size() && board[x][y] > board[r][c]) path = max(path, 1 + dfs(x, y, board, memo));
    }
    memo[r][c] = path;
    return path;
}

int longest_increasing_path(vector<vector<int>> board) {
    int res = INT_MIN;
    int n = board.size();
    int m = board[0].size();
    vector<vector<int>> memo(n, vector<int>(m, 0));
    for (int i = 0; i < board.size(); i++) {
        for (int  j = 0; j < board[0].size(); j++) {
            res = max(res, dfs(i, j, board, memo));
        }
    }
    return res;
}

int main(int argc, char **argv) {
    vector<vector<int>> board = {{1, 5, 8}, {3, 4, 4}, {7, 9, 2}};
    cout << longest_increasing_path(board) << "\n";
    return 0;
}
