#include <iostream>
#include <set>
#include <vector>

using namespace std;

void dfs(int r, set<int> diag, set<int> anti_diag, set<int> cols, int n, int& res) {
    if (r == n) {
        res++;
        return;
    }
    
    for (int c = 0; c < n; c++) {
        int curr_diag = r - c;
        int curr_anti_diag = r + c;
        if (diag.count(curr_diag) || anti_diag.count(curr_anti_diag) || cols.count(c)) continue;
        cols.insert(c);
        diag.insert(curr_diag);
        anti_diag.insert(curr_anti_diag);
        dfs(r + 1, diag, anti_diag, cols, n, res);
        cols.erase(c);
        diag.erase(curr_diag);
        anti_diag.erase(curr_anti_diag);
    }
}

int n_queens(int n) {
    int res = 0;
    dfs(0, {}, {}, {}, n, res);
    return res;
}

int main(int argc, char** argv) {
    int n = 4;
    cout << n_queens(n) << "\n";
    return 0;
}
