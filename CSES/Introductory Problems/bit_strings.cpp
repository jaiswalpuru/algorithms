#include <iostream>

using namespace std;

typedef long long ll;

ll mod = 1000000007;

ll bit_strings(ll n) {
    if (n == 0){
        return 1;
    }
    if (n%2==0){
        ll half = bit_strings(n/2);
        return (half*half)%mod;
    }else{
        ll half = bit_strings(n/2);
        return (2*half*half)%mod;
    }
}

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    ll n;
    cin>>n;
    cout<<bit_strings(n)<<"\n";
}
