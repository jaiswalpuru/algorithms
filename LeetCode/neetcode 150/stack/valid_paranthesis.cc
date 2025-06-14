class Solution {
public:
    bool isValid(string s) {
        int i = 0;
        stack<char> st;
        while (i < s.size()) {
            if (s[i] == '(' || s[i] == '{' || s[i] == '[') {
                st.push(s[i]);
            } else {
                if (st.empty()) return false;
                if (s[i] == ')' && st.top() != '(') return false;
                if (s[i] == '}' && st.top() != '{') return false;
                if (s[i] == ']' && st.top() != '[') return false;
                st.pop();
            }
            i++;
        }
        return st.empty();
    }
};
