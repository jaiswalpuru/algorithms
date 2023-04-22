#include <iostream>

bool ugly_number(int n) {
    if (n==0) {
        return false;
    }
    while(!(n%2)) n = n/2;
    while(!(n%3)) n = n/3;
    while(!(n%5)) n = n/5;
    return n == 1 ? true : false;
}

int main() {
    std::cout<<ugly_number(6)<<"\n";
}
