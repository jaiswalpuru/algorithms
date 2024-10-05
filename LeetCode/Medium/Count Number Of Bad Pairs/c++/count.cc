class Solution {
public:
    long long countBadPairs(vector<int>& nums) {
        long long res = 0;
        unordered_map<int, long> cnt_map;
        long total_pairs = ((long)(nums.size() * (nums.size() - 1))) / 2;

        for (int i = 0; i < nums.size(); i++) {
            cnt_map[nums[i] - i]++;
        }

        for (auto p : cnt_map) {
            if (p.second > 1) {
                res += (p.second * (p.second -1)) / 2;
            }
        }
        return total_pairs - res;
    }
};
