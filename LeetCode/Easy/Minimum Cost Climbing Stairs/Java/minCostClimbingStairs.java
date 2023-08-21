class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int[] dp = new int[cost.length];
        Arrays.fill(dp, -1);
        return Math.min(recur(cost, 0, dp), recur(cost, 1, dp));
    }

    private int recur(int[] cost, int start, int[] dp) {
        if (start >= cost.length) return 0;
        if (dp[start] != -1) return dp[start];
        int one = cost[start] + recur(cost, start+1, dp);
        int two = cost[start] + recur(cost, start+2, dp);
        dp[start] = Math.min(one, two);
        return dp[start];
    }
}
