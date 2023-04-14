class Solution {
    public int longestPalindromeSubseq(String s) {
        int n = s.length();
        int[][] memo = new int[n][n];
        return recur(0, n-1, s, memo);
    }

    private int recur(int l, int r, String s, int[][] memo) {
        if (l > r) return 0;
        if (l == r) return 1;

        if (memo[l][r] != 0) return memo[l][r];

        if (s.charAt(l) == s.charAt(r)) {
            memo[l][r] = 2 + recur(l+1, r-1, s, memo);
        } else {
            memo[l][r] = Math.max(recur(l+1, r, s, memo), recur(l, r-1, s, memo));
        }

        return memo[l][r];
    }
}
