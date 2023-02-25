class Solution {
    public int maxProfit(int[] prices) {
        int buy = 0;
        int profit = 0;
        int size = prices.length;

        for (int i=0; i<size; i++) {
            if (prices[buy] > prices[i]) {
                buy = i;
            }
            if (prices[i] > prices[buy]) {
                profit = Math.max(profit, prices[i]-prices[buy]);
            }
        }

        return profit;
    }
}
