/**
 * brute force
 */

class Solution {
    public int stoneGameII(int[] piles) {
        return dfs(0, 1, piles, true);
    }

    private int dfs(int k, int m, int[] piles, boolean alice) {
        if (k == piles.length) return 0;

        int res = alice ? 0 : Integer.MAX_VALUE;
        int total = 0;
        for (int x=1; x<=2*m; x++) {
            if (k+x > piles.length) break;
            total += piles[k+x-1];
            if (alice)
                res = Math.max(res, total+dfs(k+x, Math.max(m, x), piles, !alice));
            else res = Math.min(res, dfs(k+x, Math.max(m, x), piles, !alice));
        }
        return res;
    }
}

/**
 * optimized
 */

class Solution {
    int[][][] dp;
    public int stoneGameII(int[] piles) {
        dp = new int[2][piles.length+1][piles.length+1];
        for (int i=0; i<2; i++) {
            for (int j=0; j<=piles.length; j++) {
                for (int k=0; k<=piles.length; k++) {
                    dp[i][j][k] = -1;
                }
            }
        }
        return dfs(0, 1, piles, 1);
    }

    private int dfs(int k, int m, int[] piles, int alice) {
        if (k == piles.length) return 0;
        if (dp[alice][k][m] != -1) return dp[alice][k][m];
        int res = alice == 1 ? 0 : 1000000;
        int total = 0;
        for (int x=1; x<=2*m; x++) {
            if (k+x > piles.length) break;
            total += piles[k+x-1];
            if (alice == 1)
                res = Math.max(res, total+dfs(k+x, Math.max(m, x), piles, 0));
            else res = Math.min(res, dfs(k+x, Math.max(m, x), piles, 1));
        }
        return dp[alice][k][m] = res;
    }
}
