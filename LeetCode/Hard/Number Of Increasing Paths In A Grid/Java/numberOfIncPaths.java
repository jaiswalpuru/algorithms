class Solution {
    int n, m;
    int[][] dp;
    int mod = (int)1e9 + 7;
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    public int countPaths(int[][] grid) {
        n = grid.length;
        m = grid[0].length;
        dp = new int[n][m];
        int ans = 0;
        for (int i=0; i<n; i++){
            for (int j=0; j<m; j++) {
                ans = (ans + dfs(grid, i, j)) % mod;
            }
        }
        return ans;
    }

    int dfs(int[][] grid, int i, int j) {
        if (dp[i][j] != 0) return dp[i][j];
        int ans = 1;
        for (int[] d : dir) {
            int prevI = d[0] + i;
            int prevJ = d[1] + j;
            if (prevI >=0 && prevI < n && prevJ >=0 && prevJ < m && grid[prevI][prevJ] < grid[i][j]) {
                ans += dfs(grid, prevI, prevJ);
                ans %= mod;
            }
        }
        dp[i][j] = ans;
        return ans;
    }
}
