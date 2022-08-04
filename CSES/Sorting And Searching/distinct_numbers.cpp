#include <iostream>
#include <set>
#include <vector>

using namespace std;

typedef vector<int> vi;
typedef set<int> si;

void distinct_numbers(vi v) {
    si s;
    for (auto it=v.begin(); it!=v.end();it++){
        s.insert(*it);
    }
    cout<<s.size()<<"\n";
}

int main() {
    int n;
    cin>>n;
    vi v(n);
    for (int i=0;i<n;i++){
        cin >>v[i];
    }
    distinct_numbers(v);
    return 0;
}
