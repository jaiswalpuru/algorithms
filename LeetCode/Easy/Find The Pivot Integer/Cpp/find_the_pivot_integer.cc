#include <iostream>

int find_pivot_integer(int n) {
    for (int i=1; i<=n; i++) {
        int first = i*(i+1)/2;
        int last = n*(n+1)/2 - first + i;
        if (first == last) return i;
    }
    return -1;
}

int main() {
    std::cout<<find_pivot_integer(8)<<"\n";
}
