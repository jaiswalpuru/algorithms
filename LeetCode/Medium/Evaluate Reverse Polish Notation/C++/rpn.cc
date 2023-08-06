class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<string> st;
        for (string str : tokens) {
            if (str == "+" || str == "-" || str == "/" || str == "*") {
                int c1 = stoi(st.top());
                st.pop();
                int c2 = stoi(st.top());
                st.pop();
                st.push(to_string(eval(c1, c2, str)));
            } else {
                st.push(str);
            }
        }
        return stoi(st.top());
    }

    int eval(int c1, int c2, string c) {
        if (c == "/") return c2/c1;
        if (c == "*") return c1*c2;
        if (c == "+") return c1+c2;
        if (c == "-") return c2-c1;
        return 0;
    }
};
