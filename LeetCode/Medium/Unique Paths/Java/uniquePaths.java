class Solution {
    int[][] dp;
    public int uniquePaths(int m, int n) {
        dp = new int[m][n];
        for (int i=0; i<m; i++)
            Arrays.fill(dp[i], -1);
        
        return recur(m-1, n-1);
    }

    private int recur(int m, int n) {
        if (m == 0 && n == 0) return 1;
        if (m < 0 || n < 0) return 0;
        if (dp[m][n] != -1) return dp[m][n];

        int up =  recur(m, n-1);
        int left = recur(m-1, n);
        dp[m][n] = up+left;
        return dp[m][n];
    }
}
