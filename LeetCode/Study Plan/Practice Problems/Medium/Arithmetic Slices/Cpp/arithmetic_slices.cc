#include <iostream>
#include <vector>

typedef std::vector<int> vi;

//--------------------brute force---------------------------
int recur(vi& v, int ind, int& res) {
    if (ind < 2) {
        return 0;
    }
    int temp = 0;
    if (v[ind]-v[ind-1] == v[ind-1]-v[ind-2]) {
        temp = 1 + recur(v, ind-1, res);
        res += temp;
    }
    else recur(v, ind-1, res);
    return temp;
}

int arithmetic_slices(vi v) {
    int res = 0;
    recur(v, v.size()-1, res);
    return res;
}
//--------------------brute force---------------------------

//--------------------memoization---------------------------
int arithmetic_slices_memo(vi v) {
    vi dp(v.size());
    int sum = 0;
    for (int i=2; i<v.size(); i++) {
        if (v[i]-v[i-1] == v[i-1]-v[i-2]){
            dp[i] = 1 + dp[i-1];
            sum += dp[i];
        }
    }
    return sum;
}

//--------------------memoization---------------------------

int main() {
    std::cout<<arithmetic_slices_memo(vi{1,2,3,8,9,10})<<"\n";
}
