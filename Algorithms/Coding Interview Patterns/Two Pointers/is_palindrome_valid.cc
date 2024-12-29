#include <iostream>
#include <string>

using namespace std;

bool is_palin(string st) {
    int i = 0;
    int j = st.length() - 1;

    while (i < j) {
        while (i < j && !isalpha(st[i])) i++;
        while (i < j && !isalpha(st[j])) j--;
        if (st[i] != st[j]) return false;
        i++;
        j--;
    }

    return true;
}


int main(int argc, char **argv) {
    cout << is_palin("racecar");
    return 0;
}
