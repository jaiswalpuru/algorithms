class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        map<string, vector<string>> visited;
        vector<vector<string>> res;
        string temp = "";
        
        for (string str : strs) {
            temp = str;
            sort(temp.begin(), temp.end());
            visited[temp].push_back(str);
        }

        for (const auto& [key, val] : visited) {
            res.push_back(val);
        }

        return res;
    }
};
