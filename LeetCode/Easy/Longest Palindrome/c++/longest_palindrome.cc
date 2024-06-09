class Solution {
public:
    int longestPalindrome(string s) {
        int len = 0;
        int odd = 0, eve = 0;
        bool is_odd_present = false;
        map<char, int> char_cnt;
        
        if (s.length() == 1) return len+1;

        for (int i = 0; i < s.length(); i++) char_cnt[s[i]]++;

        for (const auto &it : char_cnt) {
            if (it.second % 2 == 0) len += it.second;
            else {
                is_odd_present = true;
                len += it.second - 1;
            }
        }

        return len + (is_odd_present ? 1 : 0);
    }
};
