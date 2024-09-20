class Solution {
public:
    int countConsistentStrings(string allowed, vector<string>& words) {
        int res = 0;
        char alpha[26];

        for (char c : allowed) alpha[c - 'a'] = 1;
        for (string s : words) {
            bool is_con = true;
            for (char c : s) {
                if (!alpha[c - 'a']) {
                    is_con = false;
                    break;
                }
            }
            if (is_con) res++;
        }

        return res;
    }
};
