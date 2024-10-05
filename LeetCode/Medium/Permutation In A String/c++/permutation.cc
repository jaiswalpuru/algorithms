class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        int s1_cnt[26] = {0};

        if (s1.length() > s2.length()) return false;
        for (char c : s1) s1_cnt[c - 'a']++;
        for (int i = 0; i <= s2.length() - s1.length(); i++) {
            int s2_cnt[26] = {0};
            for (int k = i; k < i + s1.length(); k++) s2_cnt[s2[k] - 'a']++;
            if (check(s1_cnt, s2_cnt)) return true;
        }

        return false;
    }

    bool check(int s1_cnt[26], int s2_cnt[26]) {
        for (int k = 0; k < 26; k++) {
            if (s1_cnt[k] != s2_cnt[k]) return false;
        }
        return true;
    }
};
