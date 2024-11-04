class Solution {
public:
    string compressedString(string word) {
        int i = 0;
        int cnt = 0;
        int ind = 0;
        string res = "";

        while (i < word.length()) {
            ind = i;
            while (i < word.length()) {
                if (word[i] == word[ind]) {
                    cnt++;
                } else {
                    break;
                }
                i++;
            }

            while (cnt) {
                if (cnt > 9) {
                    res += to_string(9) + word[ind];
                    cnt -= 9;
                } else {
                    res += to_string(cnt) + word[ind];
                    cnt = 0;
                }
            }
            cnt = 0;
        }

        while (cnt) {
            if (cnt > 9) {
                res += to_string(9) + word[ind];
                cnt -= 9;
            } else {
                res += to_string(cnt) + word[ind];
                cnt = 0;
            }
        }

        return res;
    }
};
