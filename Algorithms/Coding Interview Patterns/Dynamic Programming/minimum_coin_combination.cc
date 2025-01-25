#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

int min_coin_com_bottom(vector<int> arr, int target) {
    vector<int> dp(target + 1, INT_MAX);
    dp[0] = 0;
    for (int i = 1; i <= target; i++) {
        for (int coin : arr) {
            if (coin <= i) {
                dp[i] = min(dp[i], 1 + dp[i - coin]);
            }
        }
    }
    return dp[target] == INT_MAX ? -1 : dp[target];
}

int dfs(vector<int> arr, int target, unordered_map<int, int>& memo) {
    if (target == 0) return 0;
    if (memo.count(target)) return memo[target];
    int min_coin = INT_MAX;

    for (int i = 0; i < arr.size(); i++) {
        if (target - arr[i] >= 0) {
            min_coin = min(min_coin, 1 + dfs(arr, target - arr[i], memo));
        }
    }
    memo[target] = min_coin;
    return min_coin;
}

int min_coin_com(vector<int> arr, int target) {
    unordered_map<int, int> memo;
    int res = dfs(arr, target, memo);
    return res < 0 ? -1 : res;
}

int main(int argc, char** argv) {
    cout << min_coin_com_bottom({1, 2, 3}, 5) << "\n";
    cout << min_coin_com_bottom({2, 4}, 5) << "\n";
    return 0;
}
