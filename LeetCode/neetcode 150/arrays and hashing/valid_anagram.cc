class Solution {
public:
    bool isAnagram(string s, string t) {
        unordered_map<char, int> hash_map;
        for (char c : s) hash_map[c]++;
        for (char c : t) {
            if (hash_map.count(c)) {
                hash_map[c]--;
                if (hash_map[c] == 0) hash_map.erase(c);
            } else {
                return false;
            }
        }
        return hash_map.size() == 0;
    }
};
