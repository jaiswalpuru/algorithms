class Solution {
public:
    char findKthBit(int n, int k) {
        string s = "0";
        int len = 1;

        for (int i = 0; i < n; i++) {
            string temp = s;
            reverse(temp.begin(), temp.end());
            s = s + '1' + invert(temp);
            if (s.length() >= k) break;
        }

        return s[k-1];
    }

    string invert(string s) {
        vector<char> v;
        for (char c : s) {
            if (c == '1') v.push_back('0');
            else v.push_back('1');
        }
        string str(v.begin(), v.end());

        return str;
    }
};
