#include <iostream>
#include <vector>

typedef std::vector<std::string> vs;
typedef unsigned int ui;

vs fizz_buzz(ui n) {
    vs v;
    for (int i=1; i<=n;i++) {
        if (i%15==0) v.push_back("FizzBuzz");
        else if (i%5==0) v.push_back("Buzz");
        else if (i%3==0) v.push_back("Fizz");
        else v.push_back(std::to_string(i));
    }
    return v;
}

int main() {
    ui n = 15;
    vs v;
    v = fizz_buzz(n);
    for (auto itr=v.begin();itr!=v.end();itr++){
        std::cout<<*itr<<" ";
    }
    std::cout<<"\n";
}
