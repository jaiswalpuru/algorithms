#include <iostream>
#include <vector>

using namespace std;

int climbing_stairs_bottom(int n) {
    vector<int> arr(n + 1);
    arr[1] = 1;
    arr[2] = 2;
    for (int i = 3; i <= n; i++) {
        arr[i] = arr[i-1] + arr[i-2];
    }
    return arr[n];
}

int _dfs(vector<int>& memo, int n) {
    if (memo[n] != -1) return memo[n];
    memo[n] = _dfs(memo, n - 1) + _dfs(memo, n - 2);
    return memo[n]; 
}

int climbing_stairs_memo(int n) {
    vector<int> memo(n + 1, -1);
    memo[0] = 0;
    memo[1] = 1;
    memo[2] = 2;
    return _dfs(memo, n);
}

int climbing_stairs_recursive(int n) {
    if (n == 1 || n == 2) return n;
    return climbing_stairs_recursive(n - 1) + climbing_stairs_recursive(n - 2);
}

int main(int argc, char** argv) {
    cout << climbing_stairs_bottom(4) << "\n";
    return 0;
}
