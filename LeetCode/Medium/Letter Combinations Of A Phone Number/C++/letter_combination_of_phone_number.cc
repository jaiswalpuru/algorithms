typedef vector<string> vs;

class Solution {
public:
    map<char, string> digit_char;
    vector<string> letterCombinations(string digits) {
        vs res;
        digit_char = {{'1', ""}, {'2', "abc"},
                    {'3', "def"}, {'4', "ghi"},
                    {'5', "jkl"}, {'6', "mno"},
                    {'7', "pqrs"}, {'8', "tuv"},
                    {'9', "wxyz"}};
        backtrack(digits, 0, res, "");
        return res;
    }

    void backtrack(string digits, int ind, vs& res, string temp) {
        if (ind == digits.length()) {
            if (temp.length() > 0) res.push_back(temp);
            return;
        }

        string st = digit_char.at(digits[ind]);
        for (int i=0; i<st.length(); i++) {
            temp += st[i];
            backtrack(digits, ind+1, res, temp);
            temp.pop_back();
        }
    }
};
