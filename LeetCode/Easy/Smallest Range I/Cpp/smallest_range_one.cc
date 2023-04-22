#include <iostream>
#include <vector>

typedef std::vector<int> vi;

int smallest_range_one(vi v, int k) {
    sort(v.begin(), v.end());
    int min_val = v[0] + k;
    int max_val = v[v.size()-1] - k;
    return std::max(0, max_val-min_val);
}

int main() {
    std::cout<<smallest_range_one(vi{1}, 0)<<"\n";
}
