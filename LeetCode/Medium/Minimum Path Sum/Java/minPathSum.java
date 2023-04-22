class Solution {
    int[][] dp;
    public int minPathSum(int[][] grid) {
        int m = grid.length-1;
        int n = grid[0].length-1;
        
        dp = new int[m+1][n+1];
        for (int[] row : dp) Arrays.fill(row, -1);
        
        return recur(m, n, grid);
    }

    public int recur(int m, int n, int[][] grid) {
        if (m == 0 && n == 0) return grid[0][0];
        if (!isSafe(m, n)) return (int)1e9;

        if (dp[m][n] != -1) return dp[m][n];

        int left = grid[m][n] + recur(m, n-1, grid);
        int up = grid[m][n] + recur(m-1, n, grid);
        dp[m][n] = Math.min(left, up);
        return dp[m][n];
    }

    public boolean isSafe(int i, int j) {
        return i >= 0 && j >= 0;
    }
}
