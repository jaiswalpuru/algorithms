class Solution {
public:
    string makeFancyString(string s) {
        string res = "";
        int i = 0;

        while (i < s.length()) {
            char c = s[i];
            int cnt = 0;
            while (c == s[i]) {
                i++;
                cnt++;
            }
            for (int j = 0; j < min(cnt, 2); j++) res += c;
        }

        if (res == "") return s;
        return res;
    }
};
