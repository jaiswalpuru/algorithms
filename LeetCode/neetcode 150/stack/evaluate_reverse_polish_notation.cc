class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<string> st;
        for (string s : tokens) {
            if (s == "+" || s == "-" || s == "*" || s == "/") {
                int v1 = stoi(st.top());
                st.pop();
                int v2 = stoi(st.top());
                st.pop();
                if (s == "+") st.push(to_string(v1 + v2));
                if (s == "-") st.push(to_string(v2 - v1));
                if (s == "*") st.push(to_string(v1 * v2));
                if (s == "/") st.push(to_string(v2 / v1));
            } else {
                st.push(s);
            }
        }
        return stoi(st.top());
    }
};
