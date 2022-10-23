#include <iostream>
#include <vector>
#include <numeric>

typedef std::vector<int> vi;

int number_of_subarrays_with_gcd_equals_k(vi& v, int k) {
    int cnt=0;
    for (int i=0;i<v.size();i++){
        int curr = 0;
        for (int j=i;j<v.size();j++){
            curr = std::gcd(curr, v[j]);
            cnt += curr==k?1:0;
        }
    }
    return cnt;
}

int main() {
    vi v{9,3,1,2,6,3};
    std::cout<<number_of_subarrays_with_gcd_equals_k(v,3)<<"\n";
}
