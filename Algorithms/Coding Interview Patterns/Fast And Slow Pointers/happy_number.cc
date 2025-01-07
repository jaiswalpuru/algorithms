#include <iostream>

using namespace std;

int get_next_num(int n) {
    int next_num = 0;
        
    while(n) {
        next_num += pow(n%10, 2);
        n = n/10;
    }

    return next_num;
}

bool happy_number(int n) {
    int slow = n;
    int fast = n;

    while (true) {
        slow = get_next_num(slow);
        fast = get_next_num(get_next_num(fast));

        if (fast == 1) return true;
        if (fast == slow) return false;
    }
    
    return false;
}

int main(int argc, char **argv) {
    cout << happy_number(23) << "\n";
    cout << happy_number(116) << "\n";
    return 0;
}
