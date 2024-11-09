class Solution {
public:
    bool validPalindrome(string s) {
        return valid_palindrome(s, 0, s.length() - 1, false);
    }

    bool valid_palindrome(string s, int i, int j, bool used) {
        while (i < j) {
            if (s[i] != s[j]) {
                if (used) return false;
                return valid_palindrome(s, i + 1, j, true) || valid_palindrome(s, i, j - 1, true);
            } else {
                i++;
                j--;
            }
        }

        return true;
    }
};
