#include <iostream>
#include <stack>

using namespace std;

int evaluate(string s) {
    stack<int> st;
    int cur_num = 0;
    int res = 0; 
    int sign = 1;

    for (char c : s) {
        if (isdigit(c)) {
            cur_num = 10 * cur_num + (c - '0');
        } else if (c == '+' || c == '-') {
            res += cur_num * sign;
            sign = c == '-' ? -1 : 1;
            cur_num = 0;
        } else if (c == '(') {
            st.push(res);
            st.push(sign);
            res = 0;
            sign = 1;
        } else if (c == ')') {
             res += sign * cur_num;
             res *= st.top();
             st.pop();
             res += st.top();
             st.pop();
        }
    }

    return res + cur_num * sign;
}

int main(int argc, char** argv) {
    string s = "18-(7+(2-4))";
    cout << evaluate(s) << "\n";
    return 0;
}
