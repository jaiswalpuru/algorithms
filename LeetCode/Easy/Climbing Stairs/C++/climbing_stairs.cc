class Solution {
public:
    int climbStairs(int n) {
        if (n <= 3) return n;
        vector<int> memo(n+1);
        memo[1]=1;memo[2]=2;memo[3]=3;
        return _recur(n, memo);
    }

    int _recur(int n, vector<int>& memo) {
        if (n <= 0) return 0;
        if (memo[n] != 0) return memo[n];
        memo[n] = _recur(n-1, memo) + _recur(n-2, memo);
        return memo[n];
    }
};
