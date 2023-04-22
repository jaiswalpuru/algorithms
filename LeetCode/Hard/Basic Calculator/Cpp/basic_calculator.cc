#include <iostream>
#include <stack>
#include <cmath>

typedef std::stack<std::string> ss;

int eval(ss &st) {
    bool f = false;
    std::string h = st.top();
    if (h=="+" || h=="-") f = true;
    if (st.empty() || f) {
        st.push("0");
    }
    int res = std::stoi(st.top());
    st.pop();
    while (!st.empty() && st.top()!=")") {
        std::string sign = st.top();
        st.pop();
        if (sign == "+") res += std::stoi(st.top());
        else res -= std::stoi(st.top());
        st.pop();
    }
    return res;
}

int basic_calculator(std::string s) {
    ss st;
    int op = 0;
    int n = 0;
    for (int i=s.size()-1; i>=0; i--) {
        if (std::isdigit(s[i])){
            op = pow(10, n)*(s[i]-'0') + op;
            n++;
        }else if (s[i] != ' '){
            if (n != 0) {
                st.push(std::to_string(op));
                n = 0;
                op = 0;
            }
            if (s[i] == '(') {
                int res = eval(st);
                st.pop();
                st.push(std::to_string(res));
            }else {
                std::string t = "";
                t += s[i];
                st.push(t);
            }
        }
    }
    if (n != 0) {
        st.push(std::to_string(op));
    }
    return eval(st);
}

int main() {
    std::cout<<basic_calculator("(5-(1+(5)))")<<"\n";
}
