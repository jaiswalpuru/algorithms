#include <iostream>
#include <stack>

using namespace std;

string removal(string s) {
    stack<char> st;
    int i = 0;

    while (i < s.length()) {
        if (st.empty() || (st.top() != s[i])) {
            st.push(s[i]);
        } else {
            st.pop();
        }
        i++;
    }

    string res = "";
    while (!st.empty()) {
        res = st.top() + res;
        st.pop();
    }
    return res;
}

int main(int argc, char** argv) {
    cout << removal("aacabba") << "\n";
    cout << removal("aaa") << "\n";
    return 0;
}
