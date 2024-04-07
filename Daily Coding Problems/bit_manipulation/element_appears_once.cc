#include <iostream>
#include <vector>

typedef std::vector<int> vi;

int find_unique_element(vi v) {
    vi res(32, 0); // assuming all integers are represented with in 32 bits
    int ans = 0;

    for (int num : v) {
        for (int i=0; i<32; i++) {
            int bit = num >> i & 1;
            res[i] += bit;
        }
    }

    for (int i=0; i<res.size(); i++) {
        if (res[i]%3 != 0) {
            ans += 1 << i;
        }
    }
    
    return ans;
}

int main() {
    vi v{6, 10, 3, 3, 3, 6, 6};
    std::cout<<find_unique_element(v)<<"\n";
    return 0;
}
