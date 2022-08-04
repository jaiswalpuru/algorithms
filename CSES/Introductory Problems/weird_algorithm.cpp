#include <iostream>

typedef long long ll;

void weird_algorithm(ll n) {
    std::cout<<n<<" ";
    while (n!=1){
        if (n%2){
            n = n*3+1;
        }else{
            n /= 2;
        }
        if (n!=1){
            std::cout<<n<<" "; 
        }else{
            std::cout<<n<<"\n";        
        }
    }
}

int main() {
    ll n;
    std::cin>>n;
    weird_algorithm(n);
    return 0;
}
