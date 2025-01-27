#include <iostream>

using namespace std;

int largest_square_area(vector<vector<int>> arr) {
    int n = arr.size();
    int m = arr[0].size();
    int max_len = 0;
    vector<vector<int>> dp(n, vector<int>(m));
    for (int i = 0; i < n; i++) {
        if (arr[i][0] == 1) {
            dp[i][0] = 1;
            max_len = 1;
        }
    }
    for (int i = 0; i < m; i++) {
        if (arr[0][i] == 1) {
            dp[0][i] = 1;
            max_len = 1;
        }
    }

    for (int i = 1; i < n; i++) {
        for (int j = 1; j < m; j++) {
            if (arr[i][j] == 1) {
                dp[i][j] = 1 + min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1]));
            }
            max_len = max(max_len, dp[i][j]);
        }
    }
    return max_len * max_len;
}
    
int main(int argc, char** argv) {
    vector<vector<int>> arr = {
        {1, 0, 1, 1, 0}, {0, 0, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 0, 1}, {1, 1, 1, 0, 1}};
    cout << largest_square_area(arr) << "\n";
}
