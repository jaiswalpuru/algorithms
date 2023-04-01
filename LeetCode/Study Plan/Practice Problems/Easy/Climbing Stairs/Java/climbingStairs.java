class Solution {
    public int climbStairs(int n) {
        if (n <= 3) return n;
        int[] memo = new int[n+1];
        Arrays.fill(memo, -1);
        memo[0] = 0;
        memo[1] = 1;
        memo[2] = 2;
        return recur(n, memo);
    }

    private int recur(int n, int[] memo) {
        if (n <= 3) return n;
        if (memo[n] != -1) return memo[n];
        memo[n] = recur(n-1, memo) + recur(n-2, memo);
        return memo[n];
    }
}
