class Solution {
public:
    void wallsAndGates(vector<vector<int>>& rooms) {
        int m = rooms.size();
        int n = rooms[0].size();
        vector<vector<bool>> visited(m, vector<bool>(n, false));
        queue<vector<int>> q;
        
        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                if (rooms[r][c] == 0) q.push({r, c, 0});
            }
        }
        vector<vector<int>> dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
        while (!q.empty()) {
            int size = q.size();
            auto curr_ele = q.front();
            int row = curr_ele[0];
            int col = curr_ele[1];
            int dis = curr_ele[2];
            q.pop();
            visited[row][col] = true;
            for (int d = 0; d < 4; d++) {
                int new_row = row + dir[d][0];
                int new_col = col + dir[d][1];
                if (is_safe(new_row, new_col, m, n, visited, rooms)) {
                    rooms[new_row][new_col] = dis + 1;
                    visited[new_row][new_col] = true;
                    q.push({new_row, new_col, dis + 1});
                }
            }
        }
    }

    bool is_safe(int row, int col, int m, int n, vector<vector<bool>>& visited, vector<vector<int>>& rooms) {
        return !(row < 0 || col < 0 || row >= m || col >= n || visited[row][col] || rooms[row][col] == -1 || rooms[row][col] == 0 || rooms[row][col] != INT_MAX);
    }
};
