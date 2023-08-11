// ----- brute force -------
class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        return recur(coins, amount);
    }

    int recur(vector<int> coins, int amt) {
        if (amt < 0) return -1;
        if (amt == 0) return 0;
        int min_val = INT_MAX;
        for (int i=0; i<coins.size(); i++) {
            int take = recur(coins, amt-coins[i]);
            if (take >= 0 && take < min_val) {
                min_val = take + 1;
            }
        }
        return min_val == INT_MAX ? -1 : min_val;
    }
};


// ---- memo ------
class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        vector<int> memo(amount+1);
        return recur(coins, amount, memo);
    }

    int recur(vector<int> coins, int amt, vector<int>& memo) {
        if (amt < 0) return -1;
        if (amt == 0) return 0;
        if (memo[amt] != 0) return memo[amt];
        int min_val = INT_MAX;
        for (int i=0; i<coins.size(); i++) {
            int take = recur(coins, amt-coins[i], memo);
            if (take >= 0 && take < min_val) {
                min_val = take + 1;
            }
        }
        memo[amt] = min_val == INT_MAX ? -1 : min_val;
        return memo[amt];
    }
};
