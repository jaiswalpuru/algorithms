#include <iostream>
#include <vector>
#include <bitset>

using namespace std;

typedef vector<int> vi;

void _recur(vi &v, int n) {
    if (n == 0) {
        v.push_back(0);
        return;
    }

    _recur(v, n-1);
    int len = v.size();
    int mask = 1 << (n-1);
    for (int i= len-1; i>=0;i--){
        v.push_back(v[i]|mask);
    }
}

int get_size(int n){
    if (n != 0) {
        return n;
    }
    return 0;
}

void gray_code(int n) {
    vi v;
    _recur(v, n);
    for (auto i:v){
        bitset<16> x(i);
        for (int i=0;i<n;i++){
            cout<<x[i];
        }
        cout<<"\n";
    }
}

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    int n;
    cin>>n;
    gray_code(n);
}
