#include <iostream>

typedef long long ll;

ll sum_of_digits(ll n) {
    ll sum = 0;
    while(n) {
        sum += n%10;
        n /= 10;
    }
    return sum;
}

ll minimum_addition_to_make_integer(ll n, int target) {
    ll cnt = 0;
    int mul = 1;
    while (sum_of_digits(n) > target) {
        cnt += mul*(10-n%10);
        mul*=10;
        n=n/10+1;
    }
    return cnt;
}

int main() {
    std::cout<<minimum_addition_to_make_integer(16, 6)<<"\n";
}
