#include <iostream>

typedef long long ll;

void increasing_array(int n) {
    ll steps = 0;
    ll prev = -1, val;

    for (int i=0; i<n; i++){
        std::cin>>val;
        if (prev>val) {
            steps += prev-val;
        }else {
            prev = val;
        }
    }
    
    std::cout<<steps<<"\n";
}

int main() {
    int n;
    std::cin>>n;
    increasing_array(n);
    return 0;
}
