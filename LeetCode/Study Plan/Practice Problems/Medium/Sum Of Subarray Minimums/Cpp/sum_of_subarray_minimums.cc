#include <iostream>
#include <vector>
#include <stack>

typedef std::vector<int> vi;
typedef std::stack<int> st;
typedef long long ll;

int sum_of_subarray_minimums(vi v) {
    int mod = 1e9 + 7;
    st s;
    ll sum = 0;
    for (int i=0; i<=v.size(); i++) {
        while(!s.empty() && (i==v.size() || v[s.top()] >= v[i])) {
            int mid = s.top();
            s.pop();
            int left = s.empty() ? -1 : s.top();
            int right = i;
            int cnt = (mid-left) * (right-mid) % mod;
            sum += ((ll)cnt*v[mid])%mod;
            sum %= mod;
        }
        s.push(i);
    }
    return (int)sum;
}

int main() {
    std::cout<<sum_of_subarray_minimums(vi{3, 1, 2, 4})<<"\n";
}
