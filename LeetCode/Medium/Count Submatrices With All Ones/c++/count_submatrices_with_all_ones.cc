class Solution {
public:
    int countSquares(vector<vector<int>>& matrix) {
        int res = 0;
        vector<vector<int>> cache(matrix.size(), vector<int>(matrix[0].size()));
        for (int i = 0; i < matrix.size(); i++) {
            for (int j = 0; j < matrix[0].size(); j++) {
                cache[i][j] = -1;
            }
        }

        for (int i = 0; i < matrix.size(); i++) {
            for (int j = 0; j < matrix[0].size(); j++) {
                res += dfs(matrix, i, j, cache);
            }
        }

        return res;
    }

    int dfs(vector<vector<int>>& matrix, int r, int c, vector<vector<int>>& cache) {
        if (r >= matrix.size() || c >= matrix[0].size() || matrix[r][c] == 0) return 0;
        
        if (cache[r][c] != -1) return cache[r][c];

        int res = 1 + min(dfs(matrix, r + 1, c, cache),
                min(dfs(matrix, r, c + 1, cache), dfs(matrix, r + 1, c + 1, cache))
                );

        return cache[r][c] = res;     
    }
};
