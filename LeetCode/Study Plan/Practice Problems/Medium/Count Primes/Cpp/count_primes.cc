#include <iostream>
#include <vector>

typedef std::vector<bool> vb; 
typedef long long ll;

vb sieve(ll n) {
    vb is_prime(n);
    for (ll i=2; i<n; i++) is_prime[i] = true;
    for (ll i=2; i<n; i++) {
        if (!is_prime[i]) continue;
        for (ll j= i*i; j<n; j+=i) {
            is_prime[j] = false;
        }
    }
    return is_prime;
}

int count_primes(int n) {
    vb is_prime = sieve(n);
    int cnt = 0;
    for (int i=2; i<n; i++) if (is_prime[i]) cnt++;
    return cnt;
}

int main() {
    std::cout<<count_primes(10)<<"\n";
}
