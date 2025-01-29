#include <iostream>
#include <string>

using namespace std;

string next_alpha_seq(string s) {
    int pivot = s.length() - 2;
    int right_succ = s.length() - 1;
    while (pivot >= 0 && s[pivot] >= s[pivot + 1]) {
        pivot--;
    }
    if (pivot == -1) {
        reverse(s.begin(), s.end());
        return s;
    }

    while (s[right_succ] <= s[pivot]) {
        right_succ--;
    }
    
    char c = s[pivot];
    s[pivot] = s[right_succ];
    s[right_succ] = c;

    reverse(s.begin() + pivot + 1, s.end());

    return s;
}

int main(int argc, char **argv) {
    cout << next_alpha_seq("abcd") << "\n";
    return 0;
}
