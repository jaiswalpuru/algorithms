class Solution {
public:
    vector<int> divisibilityArray(string word, int m) {
        long long val = 0;
        vector<int> div(word.length());

        for (int i = 0; i < word.length(); i++) {
            val = ((val * 10) + (word[i] - '0')) % m;
            if (val % m == 0) div[i] = 1;
            else div[i] = 0;
        }

        return div;
    }
};
