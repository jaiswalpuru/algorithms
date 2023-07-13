class Solution {
    public int maxProfit(int[] prices) {
        int minInd = 0;
        int profit = 0;
        for (int i=0; i<prices.length; i++) {
            if (prices[minInd] > prices[i]) minInd = i;
            else profit = Math.max(profit, prices[i]-prices[minInd]);
        }
        return profit;
    }
}
