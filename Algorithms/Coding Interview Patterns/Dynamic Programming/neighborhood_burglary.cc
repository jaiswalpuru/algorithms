#include <iostream>
#include <vector>

using namespace std;

int bottom_up(vector<int> arr) {
    vector<int> dp(arr.size(), 0);
    dp[0] = arr[0];
    dp[1] = arr[1];

    for (int i = 2; i < arr.size(); i++) {
        dp[i] = max(arr[i] + dp[i-2], dp[i-1]);
    }
    return dp[arr.size()-1];
}

int dfs(vector<int> arr, int ind, vector<int>& memo) {
    if (ind >= arr.size()) {
        return 0;
    }
    if (memo[ind] != -1) return memo[ind];

    int dnt_take = dfs(arr, ind + 1, memo);
    int take = arr[ind] + dfs(arr, ind + 2, memo);
    return memo[ind] = max(take, dnt_take);
}

int max_cash(vector<int> arr) {
    vector<int> memo(arr.size(), -1);
    return dfs(arr, 0, memo);
}

int main(int argc, char** argv) {
    cout << bottom_up({200, 300, 200, 50}) << "\n";
    return 0;
}
