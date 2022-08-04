#include <iostream>

using namespace std;

typedef long long ll;

ll trailing_zeros(ll n) {
    ll cnt = 0, val = 5;
    while (n/val!=0){
        cnt += (n/val);
        val *= 5;
    }
    return cnt;
}

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    ll n;
    cin>>n;
    cout<<trailing_zeros(n)<<"\n";
}
