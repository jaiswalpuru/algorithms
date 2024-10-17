class Solution {
public:
    int maximumSwap(int num) {
        string s = to_string(num);

        int n = s.length();
        int max_num = num;

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                swap(s[i], s[j]);
                max_num = max(max_num, stoi(s));
                swap(s[i], s[j]);
            }
        }

        return max_num;
    }
};
