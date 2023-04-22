#include <iostream>
#include <map>
#include <vector>

typedef std::vector<int> vi;
typedef std::map<int, int> mi;

vi set_mismatch(vi &v) {
    vi res(2);
    mi m;
    size_t n = v.size();
    for (int i=0; i<n;i++){
        m[v[i]]++;
    }
    for (int i=1; i<=n;i++){
        if (m[i] >= 2) {
            res[0] = i;
        }
        if (m[i]== 0) {
            res[1] = i;
        }        
    }
    return res;
}

int main() {
    vi v {1, 2, 2, 4};
    vi res = set_mismatch(v);
    for (auto itr = res.begin(); itr!=res.end(); itr++){
        std::cout<<*itr<<" ";
    }
    std::cout<<"\n";
}
