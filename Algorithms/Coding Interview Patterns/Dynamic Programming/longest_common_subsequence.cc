#include <iostream>

using namespace std;

int bottom_up(string s1, string s2) {
    vector<vector<int>> dp(s1.length() + 1, vector<int>(s2.length() + 1, 0));
    for (int i = 1; i <= s1.length(); i++) {
        for (int j = 1; j <= s2.length(); j++) {
            if (s1[i - 1] == s2[j - 1]) {
                dp[i][j] = 1 + dp[i - 1][j -1];
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
            }
        }
    }
    return dp[s1.length()][s2.length()];
}

int dfs(string s1, string s2, int i, int j, vector<vector<int>> memo) {
    if (i >= s1.length() || j >= s2.length()) return 0;
    if (memo[i][j] != -1) return memo[i][j];
    if (s1[i] == s2[j]) return memo[i][j] = 1 + dfs(s1, s2, i + 1, j + 1, memo);
    return memo[i][j] = max(dfs(s1, s2, i + 1, j, memo), dfs(s1, s2, i, j + 1, memo));
}

int longest_common(string s1, string s2) {
    vector<vector<int>> memo(s1.length(), vector<int>(s2.length(), -1));
    return dfs(s1, s2, 0, 0, memo);
}

int main(int argc, char** argv) {
    cout << bottom_up("acabac", "aebab") << "\n";
    return 0;
}
