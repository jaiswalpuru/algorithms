#include <iostream>
#include <vector>

typedef std::vector<int> vi;

int frequency_of_most_frequent_element(vi &v, int k) {
    size_t sum=0, res=1, l=0, n=v.size();
    std::sort(v.begin(), v.end());
    for (int r=0;r<n;r++) {
        sum += v[r];
        while (v[r]*(r-l+1)-sum>k) {
            sum -= v[l];
            l++;
        }
        res = std::max(res, r-l+1);
    }
    return res;
}

int main() {
    vi v{1,2,4};
    std::cout<<frequency_of_most_frequent_element(v, 5)<<"\n";
}
