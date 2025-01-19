#include <iostream>
#include <vector>
#include <set>

using namespace std;

vector<vector<int>> dir = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

void dfs(vector<vector<int>> board, int r, int c, vector<vector<bool>>& visited) {
    if (visited[r][c]) return;
    visited[r][c] = true;

    for (int i = 0; i < dir.size(); i++) {
        int x = r + dir[i][0];
        int y = c + dir[i][1];
        if (x >= 0 && y >= 0 && x < board.size() && y < board[0].size() && board[x][y] == 1 && !visited[x][y]) dfs(board, x, y, visited);
    }
}

int count_islands(vector<vector<int>> board) {
    int islands = 0;
    int n = board.size();
    int m = board[0].size();
    vector<vector<bool>> visited(n, vector<bool>(m, false));

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            if (board[i][j] == 1 && !visited[i][j]) {
                islands++;
                dfs(board, i, j, visited);
            }
        }
    }
    return islands;
}

int main(int argc, char** argv) {
    vector<vector<int>> board = {{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 1, 1}, {0, 0, 0, 1}};
    cout << count_islands(board) << "\n";
    return 0;
}
