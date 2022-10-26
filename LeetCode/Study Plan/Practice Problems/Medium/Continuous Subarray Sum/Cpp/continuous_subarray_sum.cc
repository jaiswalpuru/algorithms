#include <iostream>
#include <vector>
#include <unordered_map>

typedef std::vector<int> vi;
typedef std::unordered_map<int, int> uii;

bool continuous_subarray_sum(vi &v, int k) {
    uii hp{{0,-1}};
    int sum = 0;
    for (int i=0;i<v.size();i++){
        sum += v[i];
        int remainder = sum % k;
        if (hp.find(remainder) == hp.end()) {
            hp[remainder] = i;
        }else {
            if (i - hp[remainder] > 1) return true;
        }
    }
    return false;
}

int main() {
    vi v{23, 2, 6, 7};
    std::cout<<continuous_subarray_sum(v, 6)<<"\n";
}
