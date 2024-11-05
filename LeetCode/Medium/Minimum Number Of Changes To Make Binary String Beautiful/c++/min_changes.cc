class Solution {
public:
    int minChanges(string s) {
        char curr_char = s[0];
        int consecutive_cnt = 0;
        int min_change_reqd = 0;

        for (int i = 0; i < s.length(); i++) {
            if (s[i] == curr_char) {
                consecutive_cnt++;
                continue;
            }

            //even case
            if (consecutive_cnt % 2 == 0) {
                consecutive_cnt = 1;
            } else {
                consecutive_cnt = 0;
                min_change_reqd++;
            }
            
            curr_char = s[i];
        }

        return min_change_reqd;
    }
};
