#include <iostream>

using namespace std;

int bottom_up(vector<int>& arr) {
    vector<int> memo(arr.size(), -1e4 - 1);
    memo[0] = arr[0];
    int max_sum = memo[0];
    for (int i = 1; i < arr.size(); i++) {
        memo[i] = max(memo[i - 1] + arr[i], arr[i]);
        max_sum = max(max_sum, memo[i]);
    }
    return max_sum;
}

void dfs(int& max_sum, vector<int> arr, int ind, int curr_sum) {
    if (ind == arr.size()) return;

    curr_sum = max(arr[ind], arr[ind] + curr_sum);
    max_sum = max(max_sum, curr_sum);
    dfs(max_sum, arr, ind + 1, curr_sum);
}

int max_sub_sum(vector<int> arr) {
    int max_sum = INT_MIN;
    dfs(max_sum, arr, 0, INT_MIN);
    return max_sum;
}

int main(int argc, char** argv) {
    cout << max_sub_sum({3, 1, -6, 2, -1, 4, -9}) << "\n";
    return 0;
}
