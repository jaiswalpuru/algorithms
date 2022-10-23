#include <iostream>
#include <numeric>
#include <vector>
#include <algorithm>

typedef unsigned long ui;
typedef std::vector<int> vi;

int greatest_common_divisor_of_array(vi &v) {
    size_t min_val=v[0], max_val=v[0];
    for (int i=0;i<v.size();i++) {
        min_val = std::min((ui)min_val, (ui)v[i]);
        max_val = std::max((ui)max_val, (ui)v[i]);
    }
    return std::gcd(min_val, max_val);
}

int main() {
    vi v{2,5,6,9,10};
    std::cout<<greatest_common_divisor_of_array(v)<<"\n";
}
