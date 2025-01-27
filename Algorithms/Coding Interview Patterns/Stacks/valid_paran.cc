#include <iostream>
#include <stack>

using namespace std;

bool valid(string s) {
    stack<char> st;
    for (char c : s) {
        if (c == '[' || c == '{' || c == '(') {
            st.push(c);
        } else {
            if (!st.empty() && ((c == ']' && st.top() == '[') || (c == '}' && st.top() == '{') || (c == ')' && st.top() == '('))) {
                st.pop();
            } else {
                return false;
            }
        }
    }
    return st.empty();
}

int main(int argc, char** argv) {
    cout << valid("([]{})") << "\n";
    cout << valid("([]{)}") << "\n";
    return 0;
}
