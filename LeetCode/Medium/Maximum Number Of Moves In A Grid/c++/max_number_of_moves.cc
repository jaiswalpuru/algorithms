class Solution {
public:
    int dir[3] = {-1, 0, 1};
    int maxMoves(vector<vector<int>>& grid) {
        int r = 0;
        int c = 0;
        int moves = 0;
        vector<vector<int>> dp(grid.size(), vector<int>(grid[0].size(), -1));
        for (int i = 0; i < grid.size(); i++) {
            moves = max(recur(grid, i, 0, dp), moves);
        }

        return moves;
    }

    int recur(vector<vector<int>>& grid, int r, int c, vector<vector<int>>& dp) {
        int m = 0;
        if (dp[r][c] != -1) return dp[r][c];

        for (int i = 0; i < 3; i++) {
            int new_r = r + dir[i];
            int new_c = c + 1;
            if (new_r >= 0 && new_c >= 0 && new_r < grid.size() && new_c <grid[0].size() &&
            grid[r][c] < grid[new_r][new_c]){
                m = max(m, 1 + recur(grid, new_r, new_c, dp));
            }
        }
        return dp[r][c] = m;
    }
};
