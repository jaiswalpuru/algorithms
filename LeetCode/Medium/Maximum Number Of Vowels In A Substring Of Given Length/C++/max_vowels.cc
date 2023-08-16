class Solution {
public:
    int maxVowels(string s, int k) {
        int i=0, max_vowel=0, curr=0;
        int j=0;
        while(j < s.length()) {
            if (j-i== k) {
                max_vowel = max(max_vowel, curr);
                if (is_vowel(s[i])) curr--;
                i++;
            }
            if (is_vowel(s[j])) curr++;
            j++;
        }
        max_vowel = max(max_vowel, curr);
        return max_vowel;
    }

    bool is_vowel(char c) {
        return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
    }
};
