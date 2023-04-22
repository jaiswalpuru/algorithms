class Solution {
    public int coinChange(int[] coins, int amount) {
        if (amount < 1) return 0;
        return recur(coins, amount, new int[amount]);
    }

    private int recur(int[] coins, int amount, int[] memo) {
        if (amount < 0) return -1;
        if (amount == 0) return 0;
        if (memo[amount-1] != 0) return memo[amount-1];

        int minVal = Integer.MAX_VALUE;
        for (int i=0; i<coins.length; i++) {
            int take = recur(coins, amount-coins[i], memo);
            if (take >= 0 && take < minVal) {
                minVal = take+1;
            }
        }
        memo[amount-1] = (minVal == Integer.MAX_VALUE) ? -1 :  minVal;
        
        return memo[amount-1];
    }
}
