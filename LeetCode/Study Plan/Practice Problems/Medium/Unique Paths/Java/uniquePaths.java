class Solution {
    int m;
    int n;
    public int uniquePaths(int m, int n) {
        this.m = m;
        this.n = n;
        int[][] memo = new int[m][n];
        for (int[] row : memo) Arrays.fill(row, -1);
        return recur(0, 0, memo);
    }

    private int recur(int i, int j, int[][] memo) {
        if (i == m-1 && j == n-1) return 1;
        if (i+1 > m || j+1> n) return 0;
        if (memo[i][j] != -1) return memo[i][j];
        int up = recur(i+1, j, memo);
        int right = recur(i, j+1, memo);
        memo[i][j] = up+right;
        return up+right;
    }
}
