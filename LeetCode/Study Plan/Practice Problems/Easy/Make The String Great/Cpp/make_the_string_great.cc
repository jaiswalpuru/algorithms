#include <iostream>
#include <stack>

typedef std::stack<char> sc;

std::string make_the_string_great(std::string st) {
    sc s;
    for (int i=0; i<st.length(); i++) {
        if (s.empty()) s.push(st[i]);
        else {
            char t = s.top();
            char curr = st[i];
            if (tolower(t) == tolower(curr)) {
                if ((t-curr)!=0) s.pop();
                else s.push(st[i]);
            } else {
                s.push(st[i]);
            }
        }
    }
    std::string res = "";
    while(!s.empty()) {
        res = s.top() + res;
        s.pop();
    }
    return res;
}

int main() {
    std::cout<<make_the_string_great("leEeetcode")<<"\n";
}
