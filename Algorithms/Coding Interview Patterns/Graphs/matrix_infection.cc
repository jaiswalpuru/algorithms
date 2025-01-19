#include <iostream>
#include <vector>
#include <queue>

using namespace std;

int matrix_infectiono(vector<vector<int>> board) {
    int n = board.size();
    int m = board[0].size();
    vector<vector<bool>> vis(n, vector<bool>(m, false));;
    queue<pair<int, int>> q;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            if (board[i][j] == 2) q.push({i, j});
        }
    }
    
    int time = 0;
    vector<vector<int>> dir = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
    while (!q.empty()) {
        int size = q.size();
        time++;
        for (int k = 0; k < size; k++) {
            auto p = q.front();
            q.pop();
            int r = p.first;
            int c = p.second;
            if (vis[r][c]) continue;
            vis[r][c] = true;
            for (int i = 0; i < 4; i++) {
                int x = r + dir[i][0];
                int y = c + dir[i][1];
                if (x >=0 && y >=0 && x < n && y < m && !vis[x][y] && board[x][y] == 1) {
                    board[x][y] = 2;
                    q.push({x, y});
                }
            }
        }
    }
    return time - 1;
}


int main(int argc, char **argv) {
    vector<vector<int>> board = {{1, 1, 1, 0}, {0, 0, 2, 1}, {0, 1, 1, 0}};
    cout << matrix_infectiono(board) << "\n";
    return 0;
}
