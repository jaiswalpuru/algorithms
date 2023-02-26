class Solution {
    int[][] dp;
    int n;
    int m;

    public int minDistance(String word1, String word2) {
        n = word1.length();
        m = word2.length();
        dp = new int[n][m];
        
        for (int i=0; i<n; i++) {
            Arrays.fill(dp[i], -1);
        }

        return editDistance(0, 0, word1, word2);
    }

    public int editDistance(int i, int j, String word1, String word2) {
        if (i == n) return m-j;
        if (j == m) return n-i;
        if (dp[i][j] != -1) return dp[i][j];

        if (word1.charAt(i) == word2.charAt(j)) dp[i][j] = editDistance(i+1, j+1, word1, word2);
        else {
            int insert = 1 + editDistance(i, j+1, word1, word2);
            int remove = 1 + editDistance(i+1, j, word1, word2);
            int replace = 1 + editDistance(i+1, j+1, word1, word2);
            dp[i][j] = Math.min(insert, Math.min(remove, replace));
        }

        return dp[i][j];
    }
}
