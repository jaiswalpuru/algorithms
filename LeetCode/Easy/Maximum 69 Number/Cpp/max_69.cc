#include <iostream>

int max_69(int  n) {
    std::string s = std::to_string(n);
    for (int i=0;i<s.size();i++) {
        if (s[i] == '6') {
            s[i] = '9';
            break;
        }
    }
    return stoi(s);
}

int main() {
    std::cout<<max_69(9669)<<"\n";
}
