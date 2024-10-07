class Solution {
public:
    int minLength(string s) {
        int cnt = 0;
        int i = 0;
        stack<char> st;

        while(i < s.length()) {
            if (s[i] == 'D') {
                if (!st.empty() && st.top() == 'C') {
                    st.pop();
                    cnt += 2;
                } else st.push(s[i]);
            } else if (s[i] == 'B') {
                if (!st.empty() && st.top() == 'A') {
                    st.pop();
                    cnt += 2;
                } else st.push(s[i]);
            } else st.push(s[i]);
            i++;
        }

        return s.length() - cnt;
    }
};
