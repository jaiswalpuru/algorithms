#include <iostream>
#include <vector>
#include <map>

#define LL long long

typedef std::vector<int> vi;
typedef std::vector<std::map<LL, int>> vm;

int arithmetic_slices_two(vi v) {
    vm cnt(v.size());
    LL res = 0;
    for (int i=0; i<v.size(); i++) {
        for (int j=0; j<i; j++) {
            LL d = (LL)v[i] - (LL)v[j];
            int sum = 0;
            if (cnt[j].find(d) != cnt[j].end()) sum = cnt[j][d];
            cnt[i][d] += sum + 1;
            res += sum;
        }
    }
    return (int)res;
}

int main () {
    std::cout<<arithmetic_slices_two(vi{1,1,2,3,4,5})<<"\n";
}
