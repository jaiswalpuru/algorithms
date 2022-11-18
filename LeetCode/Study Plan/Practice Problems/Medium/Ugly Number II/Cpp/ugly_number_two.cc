#include <iostream>

int ugly_number_two(int n) {
    int *dp = new int[n];
    dp[0] = 1;
    int count = 1;
    int ind_2 = 0, ind_3 = 0, ind_5 = 0;
    while(count < n) {
        int v1 = dp[ind_2]*2, v2 = dp[ind_3]*3, v3 = dp[ind_5]*5;
        int val = std::min(v1, std::min(v2, v3));
        dp[count] = val;
        count++;
        if (v1 == val) {
            ind_2++;
        }
        if (v2 == val) {
            ind_3++;
        }
        if (v3 == val) {
            ind_5++;
        }
    }
    return dp[n-1];
}

int main() {
    std::cout<<ugly_number_two(10)<<"\n";
}
    
