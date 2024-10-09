class Solution {
public:
    string minRemoveToMakeValid(string s) {
        stack<char> st;
        string res = "";
        int open = 0;

        for (int i = 0; i < s.length(); i++) {
            if (isalpha(s[i])) {
                st.push(s[i]);
            } else {
                if (s[i] == '(') {
                    open++;
                    st.push(s[i]);
                } else {
                    if (open > 0) {
                        open--;
                        st.push(s[i]);
                    }
                }
            }
        }
        
        while (!st.empty()) {
            if (open > 0 && st.top() == '(') {
                open--;
                st.pop();
                continue;
            }
            res += st.top();
            st.pop();
        }
        reverse(res.begin(), res.end());

        return res;
    }
};
