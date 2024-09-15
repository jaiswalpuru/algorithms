class Solution {
public:
    int getLucky(string s, int k) {
        int transformed = 0;
        string num = "";

        for (int i = 0; i < s.length(); i++) num += to_string(s[i] - 97 + 1);
        while (k-- > 0) {
            for (int i = 0; i < num.length(); i++) {
                transformed += num[i]-'0';
            }
            cout<<transformed;
            num = to_string(transformed);
            transformed = 0;
        }

        return stoi(num);
    }
};
