class Solution {
public:
    vector<string> uncommonFromSentences(string s1, string s2) {
        map<string, int> s_cnt;
        vector<string> res;
        stringstream ss1(s1 + " " + s2);
        string word;

        while (ss1 >> word) s_cnt[word]++;
        for (auto& [key, val] : s_cnt) if (val == 1) res.push_back(key);

        return res;
    }
};
