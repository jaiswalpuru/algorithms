class Solution {
public:
    string minWindow(string s, string t) {
        unordered_map<char, int> t_cnt;
        unordered_map<char, int> s_cnt;

        if (s == "" || t == "") return "";

        int l = 0;
        int r = 0;
        int n = s.length();
        int formed = 0;
        int ans[3] = {-1, 0, 0};

        for (char c : t) t_cnt[c]++;
        int reqd = t_cnt.size();

        while (r < n) {
            s_cnt[s[r]]++;

            if (s_cnt[s[r]] == t_cnt[s[r]]) formed++;
            
            while (l <= r && formed == reqd) {
                if (ans[0] == -1 || (r - l + 1) < ans[0]) {
                    ans[0] = r - l + 1;
                    ans[1] = l;
                    ans[2] = r;
                }
                s_cnt[s[l]]--;
                
                if (s_cnt[s[l]] < t_cnt[s[l]]) formed--;
                l++;
            }
            r++;
        }
        
        if (ans[0] == -1) return "";
        return s.substr(ans[1], ans[2] - ans[1] + 1);
    }
};
