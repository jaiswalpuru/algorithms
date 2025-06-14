class Codec {
public:

    // Encodes a list of strings to a single string.
    string encode(vector<string>& strs) {
        string res = "";
        for (const string& s : strs) {
            for (const char c : s) {
                if (c == '/') {
                    res += "//";
                } else {
                    res += c;
                }
            }
            res += "/:";
        }
        return res;
    }

    // Decodes a single string to a list of strings.
    vector<string> decode(string s) {
        vector<string> res;
        string temp = "";
        for (size_t  i = 0; i < s.size(); i++) {
            if (i < s.size() - 1 && s[i] == '/' && s[i + 1] == ':') {
                res.push_back(temp);
                temp.clear();
                i++;   
            } else if (i < s.size() - 1 && s[i] == '/' && s[i + 1] == '/') {
                temp += '/';
                i++;
            } else {
                temp += s[i];
            }
        }
        return res;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.decode(codec.encode(strs));
