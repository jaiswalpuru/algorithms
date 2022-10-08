#include <iostream>

short count_bits(unsigned int x) {
    short num_bits = 0;
    while(x) {
        num_bits += x&1;
        x>>=1; //shift right by 1
    }
    return num_bits;
}

int main() {
    std::cout<<count_bits(10)<<"\n";
}
