class Solution {
    int[][][] memo;
    public int maxProfit(int[] prices) {
        memo = new int[prices.length][2][3];
        for (int i=0; i<prices.length; i++) {
            for (int j=0; j<2; j++) {
                Arrays.fill(memo[i][j], -1);
            }
        }
        return recur(prices, 0, 2, 1);
    }

    private int recur(int[] prices, int start, int k, int buy) {
        if (start == prices.length || k == 0) return 0;
        if (memo[start][buy][k] != -1) return memo[start][buy][k];
        if (buy == 1) {
            int b = recur(prices, start+1, k, 0) - prices[start];
            int nB = recur(prices, start+1, k, 1);
            memo[start][buy][k] = Math.max(b, nB);
        } else {
            int s = recur(prices, start+1, k-1, 1) + prices[start];
            int nS = recur(prices, start+1, k, 0);
            memo[start][buy][k] = Math.max(s, nS);
        }
        return memo[start][buy][k];
    }
}
