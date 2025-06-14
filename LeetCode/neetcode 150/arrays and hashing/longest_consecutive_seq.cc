class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        unordered_set<int> hash_map(nums.begin(), nums.end());
        int max_cnt = 0;
        for (int num : hash_map) {
            if (!hash_map.count(num - 1)) {
                int curr = 1;
                while (hash_map.count(num + 1)) {
                    curr++;
                    num = num + 1;
                }
                max_cnt = max(max_cnt, curr);
            }
        }
        return max_cnt;
    }
};
