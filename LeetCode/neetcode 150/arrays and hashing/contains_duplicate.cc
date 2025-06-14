class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        unordered_map<int, int> hash_map;
        for (int num : nums) {
            hash_map[num]++;
            if (hash_map[num] > 1) return true;
        }
        return false;
    }
};
