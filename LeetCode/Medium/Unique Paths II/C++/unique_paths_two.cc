class Solution {
public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
        vector<vector<int>> memo(obstacleGrid.size(), vector<int>(obstacleGrid[0].size(), 0));
        return recur(obstacleGrid, 0, 0, memo);
    }

    int recur(vector<vector<int>>& obs, int i, int j, vector<vector<int>>& memo) {
        if (i >= obs.size() || j >= obs[0].size()) return 0;
        if (obs[i][j] == 1) return 0;
        if (memo[i][j] != 0) return memo[i][j];
        if (i == obs.size()-1 && j == obs[0].size()-1) return 1;
        int right = recur(obs, i, j+1, memo);
        int down = recur(obs, i+1, j, memo);
        return memo[i][j] = right + down;
    }
};
