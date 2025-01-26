#include <iostream>
#include <vector>

using namespace std;

int bottom_up(int cap, vector<int> wt, vector<int> val) {
    vector<vector<int>> dp(wt.size() + 1, vector<int>(cap + 1, 0));

    for (int i = wt.size() - 1; i >= 0; i--) {
        for (int j = 1; j <= cap; j++) {
            if (wt[i] <= j) {
                dp[i][j] = max(val[i] + dp[i + 1][j - wt[i]], dp[i + 1][j]);
            } else {
                dp[i][j] = dp[i + 1][j];
            }
        }
    }
    return dp[0][cap];
}

int dfs(int ind, int cap, vector<int> wt, vector<int> val, vector<vector<int>>& memo) {
    if (ind < 0) return 0;
    if (memo[ind][cap] != -1) {
        return memo[ind][cap];
    }

    if (cap - wt[ind] >= 0) {
        memo[ind][cap] = max(val[ind] + dfs(ind - 1, cap - wt[ind], wt, val, memo), dfs(ind - 1, cap, wt, val, memo));
    } else {
        memo[ind][cap] = dfs(ind - 1, cap, wt, val, memo);
    }
    return memo[ind][cap];
}


int knapsack(int cap, vector<int> wt, vector<int> val) {
    int max_val = 0;
    vector<vector<int>> memo(wt.size(), vector<int>(cap + 1, -1));
    return dfs(wt.size() - 1, cap, wt, val, memo);
}

int main(int argc, char** argv) {
     cout << bottom_up(7, {5, 3, 4, 1}, {70, 50, 40, 10}) << "\n";
     return 0;
}
