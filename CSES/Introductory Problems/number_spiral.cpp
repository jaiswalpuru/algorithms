#include <iostream>

using namespace std;

typedef long long ll;

ll calc_value(ll x, ll y) {
    ll res;
    if (x < y){
        if (y%2 == 1){
            res = y*y - (x-1);
        }else{
            res = ((y-1)*(y-1)+1) + (x-1);
        }
    }else{
        if (x%2 == 0){
            res = x*x - (y-1);
        }else{
            res = ((x-1)*(x-1)+1) + (y-1);
        }
    }
    return res;
}

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    int t;
    cin>>t;
    ll x,y;
    for (int i=0;i<t;i++){
        cin>>x>>y;
        cout<<calc_value(x,y)<<"\n";
    }
    return 0;
}
