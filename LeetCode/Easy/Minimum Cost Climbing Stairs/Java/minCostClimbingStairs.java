class Solution {
    public int minCostClimbingStairs(int[] cost) {
        int[] dp1 = new int[cost.length];
        int[] dp2 = new int[cost.length];
        Arrays.fill(dp1, -1);
        Arrays.fill(dp2, -1);
        return Math.min(recur(cost, 0, dp1), recur(cost, 1, dp2));
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
