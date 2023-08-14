class Solution {
public:
    int compress(vector<char>& chars) {
        string res = "";
        for (int i=0; i<chars.size();) {
            res += chars[i];
            int j=i+1;
            while(j < chars.size() && chars[j] == chars[i]) j++;
            if (j-i > 1) res += to_string(j-i);
            i=j;
        }
        for (int i=0; i<res.length(); i++) {
            chars[i] = res[i];
        }
        return res.length();
    }
};
