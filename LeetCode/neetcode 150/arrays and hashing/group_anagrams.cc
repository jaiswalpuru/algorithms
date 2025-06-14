class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> res;
        unordered_map<string, vector<string>> hash_map;
        for (string s : strs) {
            string temp = s;
            sort(temp.begin(), temp.end());
            hash_map[temp].push_back(s);
        }
        for (auto [k, v] : hash_map) {
            res.push_back(v);
        }
        return res;
    }
};
