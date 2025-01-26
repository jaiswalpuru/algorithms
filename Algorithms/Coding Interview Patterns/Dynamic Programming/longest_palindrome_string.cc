#include <iostream>

using namespace std;

string bottom_up(string s) {
    int n = s.length();
    if (n == 0) return "";
    int max_len = 1;
    int st_ind = 0;
    vector<vector<bool>> dp(n, vector<bool>(n, false));

    for (int i = 0; i < n; i++) dp[i][i] = true;
    for (int i = 0; i < n - 1; i++) {
        if (s[i] == s[i+1]) {
            max_len = 2;
            st_ind = i;
            dp[i][i+1] = true;
        }
    }

    for (int sub_len = 3; sub_len < n + 1; sub_len++) {
        for (int i = 0; i < n - sub_len + 1; i++) {
            int j = i + sub_len - 1;
            if (s[i] == s[j] && dp[i + 1][j - 1]) {
                dp[i][j] = true;
                max_len = sub_len;
                st_ind = i;
            }
        }
    }
    return s.substr(st_ind, max_len);
}


int is_palindrome(int i, int j, vector<vector<int>>& memo, string s) {
    if (i >= j) return 1;
    
    if (memo[i][j] != -1) return memo[i][j];

    if (s[i] == s[j]) return memo[i][j] = is_palindrome(i + 1, j - 1, memo, s);
    return memo[i][j] = 0;
}

string longest_palindrome(string s) {
    int n = s.length();
    vector<vector<int>> memo(n, vector<int>(n, -1));
    int max_len = 1;
    int start_ind = 0;

    for (int i = 0; i < n; i++) {
        for (int j = i; j < n; j++) {
            if (is_palindrome(i, j, memo, s)) {
                if (max_len < (j - i + 1)) {
                    max_len = j - i + 1;
                    start_ind = i;
                }
            }
        }
    }
    return s.substr(start_ind, max_len);
}

int main(int argc, char** argv) {
    cout << bottom_up("abccbaba") << "\n";
    return 0;
}
