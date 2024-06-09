class Solution {
public:
    vector<string> commonChars(vector<string>& words) {
        vector<string> res;
        map<char, bool> vis;
        map<int, map<char, int>> char_cnt;

        for (int i = 0; i < words.size(); i++) {
            for (int j = 0; j < words[i].length(); j++) {
                char_cnt[i][words[i][j]]++;
            }
        }
        
        for (int i = 0; i < words.size(); i++) {
            for (int j = 0; j < words[i].length(); j++) {
                char c = words[i][j];
                if (vis[c]) continue;
                vis[c] = true;
                int times = char_cnt[i][c];
                for (int j = 0; j < words.size(); j++) {
                    if (i == j) continue;
                    times = min(times, char_cnt[j][c]);
                }
                string st(1, c);
                for (int j = 0; j < times; j++) {
                    res.push_back(st);
                }
            }
        }

        return res;
    }
};
