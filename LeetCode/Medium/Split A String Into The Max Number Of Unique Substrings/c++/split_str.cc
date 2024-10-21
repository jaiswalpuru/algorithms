class Solution {
public:
    int maxUniqueSplit(string s) {
        unordered_set<string> visited;
        return backtrack(s, 0, visited);
    }

    int backtrack(const string& s, int start, unordered_set<string>& visited) {
        if (start == s.length()) return 0;
        int cnt = 0;
        
        for (int j = start + 1; j <= s.length(); j++) {
            string sub = s.substr(start, j - start);
            if (visited.find(sub) == visited.end()) {
                visited.insert(sub);
                cnt = max(cnt, 1 + backtrack(s, j, visited));
                visited.erase(sub);
            }
        }

        return cnt;
    }
};
