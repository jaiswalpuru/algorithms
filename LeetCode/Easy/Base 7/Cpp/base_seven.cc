#include <iostream>

std::string base_seven(int n) {
    bool neg = false;
    if (n<0) {
        n *= -1;
        neg = true;
    }
    if (n == 0 ) {
        return "0";
    }
    std::string s = "";
    while (n) {
        s = std::to_string(n%7) + s;
        n = n/7;
    }
    return neg ? "-" + s : s;
}

int main () {
    int n = 100;
    std::cout<<base_seven(n)<<"\n";
}
