class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        int sum = 0;
        int res = 0;
        unordered_map<int, int> cnt_map;
        cnt_map[0] = 1;

        for (int i = 0; i < nums.size(); i++) {
            sum += nums[i];
            res += cnt_map[sum-k];
            cnt_map[sum]++;
        }

        return res;
    }
};
