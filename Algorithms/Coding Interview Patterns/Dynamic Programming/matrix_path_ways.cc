#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

int matrix_path(int m, int n) {
    vector<vector<int>> dp(m, vector<int>(n));
    for (int i = 0; i < m; i++) dp[i][0] = 1;
    for (int i = 0; i < n; i++) dp[0][i] = 1;

    for (int i = 1; i < m; i++) {
        for (int j = 1; j < n; j++) {
            dp[i][j] = dp[i-1][j] + dp[i][j-1];
        }
    }
    return dp[m-1][n-1];
}

int dfs(int i, int j, int m, int n, vector<vector<int>>& dp) {
    if (i == m - 1 && j == n - 1) return 1;
    if (i >=m || j >= n) return 0;
    if (dp[i][j] != -1) return dp[i][j];

    int down = dfs(i + 1, j, m, n, dp);
    int right = dfs(i, j + 1, m, n, dp);
    dp[i][j] = down + right;
    return dp[i][j];
}

int matrix_paths(int m, int n) {
    vector<vector<int>> dp(m, vector<int>(n, -1));
    return dfs(0, 0, m, n, dp);
}

int main(int argc, char** argv) {
    cout << matrix_path(3, 3) << "\n";
    return 0;
}
