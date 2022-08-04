#include <iostream>
#include <vector>

using namespace std;

typedef long long ll;
typedef vector<ll> vl;
typedef vector<vl> vll;

void is_possible(ll n) {
    ll sum = (ll)n*(n+1)/2;
    if (sum%2 != 0) {
        cout<<"NO\n";
        return;
    }
    vl v1, v2;
    ll temp = n;
    if (n%2){
        n = n-1;
    }
    ll i = 1, j=n;
    bool flag = true;
    while (i<j) {
        if (flag){
            v1.push_back(i);
            v1.push_back(j);
        }else{
            v2.push_back(i);
            v2.push_back(j);
        }
        i++;
        j--;
        flag = !flag;
    }
    if (temp%2) {
        ll sum1 =0, sum2=0;
        for (int i=0;i<v1.size();i++){
            sum1+=v1[i];
        }
        for (int i=0;i<v2.size();i++){
            sum2+=v2[i];
        }
        if (sum1!=sum/2){
            v1.push_back(temp);
        }else{
            v2.push_back(temp);
        }
    }
    cout<<"YES\n";
    cout<<v1.size()<<"\n";
    for (int i=0;i<v1.size();i++){
        cout <<v1[i]<<" ";
    }
    cout<<"\n"<<v2.size()<<"\n";
    for (int i=0;i<v2.size();i++){
        cout<<v2[i]<<" ";
    }
    cout<<"\n";
}

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    ll n;
    cin>>n;
    is_possible(n);
    return 0;
}
