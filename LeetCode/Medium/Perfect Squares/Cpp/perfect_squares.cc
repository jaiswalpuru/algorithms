#include <iostream>
#include <vector>
#include <map>
#include <cmath>

typedef std::vector<int> vi;
typedef std::map<int, int> mi;
typedef std::map<int, bool> mb;

//------------------memo still TLE-------------------------
int min_cnt(vi& sq, mb& m, mi& visited, int n) {
    if (m.find(n) != m.end()) return 1;
    if (visited.find(n) != visited.end()) return visited[n];
    int min_val = INT_MAX;
    for (int i=1; i<sq.size(); i++) {
        if (n < sq[i]) break;
        int val = 1 + min_cnt(sq, m, visited, n-sq[i]);
        min_val = std::min(min_val, val);
        visited[n] = min_val;
    }
    return min_val;
}

int perfect_squares(int n) {
    vi sq(sqrt(n)+1);
    mb m;
    mi visited;
    for (int i=1; i<=sqrt(n); i++) {
        sq[i] = i*i;
        m[i*i] = true;
    }
    return min_cnt(sq, m, visited, n);
}
//------------------memo still TLE-------------------------

//------------------efficient approach-------------------------
int perfect_squares_eff(int n) {
    vi dp(n+1, INT_MAX);
    dp[0] = 0;
    int max_ind = (int)sqrt(n) + 1;
    vi sq(max_ind);
    for (int i=1; i<max_ind; i++) sq[i] = i*i;
    for (int i=1; i<=n; i++) {
        for (int j=1; j<max_ind; j++) {
            if (i<sq[j]) break;
            dp[i] = std::min(dp[i], dp[i-sq[j]]+1);
        }
    }
    return dp[n];
}
//------------------efficient approach-------------------------

int main () {
    std::cout<<perfect_squares_eff(12)<<"\n";
}
