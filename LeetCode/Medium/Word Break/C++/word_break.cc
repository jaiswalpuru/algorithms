class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        set<string> words(wordDict.begin(), wordDict.end());
        bool seen[s.length()+1];
        for (int i=0; i<s.length()+1; i++) seen[i] = false;
        queue<int> q;
        q.push(0);
        while(!q.empty()) {
            int start = q.front();
            q.pop();
            if (start == s.length()) return true;
            for (int end=start+1; end<=s.length(); end++) {
                if (seen[end]) continue;
                string str = s.substr(start, end-start);
                if (words.find(str) != words.end()) {
                    q.push(end);
                    seen[end] = true;
                }
            }
        }
        return false;
    }
};
