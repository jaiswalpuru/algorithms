class Solution {
    int[][] memo;
    public int maxProfit(int[] prices) {
        memo = new int[prices.length][2];
        for (int i=0; i<prices.length; i++) Arrays.fill(memo[i], -1);
        return recur(prices, 0, 1);
    }

    private int recur(int[] prices, int start, int buy) {
        if (start == prices.length) {
            return 0;
        }
        if (memo[start][buy] != -1) return memo[start][buy];
        int profit = Integer.MIN_VALUE;
        if (buy == 1) {
            int b = recur(prices, start+1, 0) - prices[start];
            int notB = recur(prices, start+1, 1);
            profit = Math.max(b, notB);
        } else {
            int s = recur(prices, start+1, 1) + prices[start];
            int notS = recur(prices, start+1, 0);
            profit = Math.max(s, notS);
        }
        memo[start][buy] = profit;
        return profit;
    }
}
