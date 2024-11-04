class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        vector<int> memo(amount + 1);
        return recur(coins, amount, memo);
    }

    int recur(vector<int>& coins, int amount, vector<int>& memo) {
        if (amount < 0) return -1;
        if (amount == 0) return 0;
        if (memo[amount] != 0) return memo[amount];
        
        int res = INT_MAX;

        for (int i = 0; i < coins.size(); i++) {
            int take = recur(coins, amount - coins[i], memo);
            if (take >= 0 && take < res) res = take + 1;
        }

        return memo[amount] = res == INT_MAX ? -1 : res;
    }
};
