class Solution {
public:
    long long countInterestingSubarrays(vector<int>& nums, int modulo, int k) {
        unordered_map<long long, long long> cnt_map;
        vector<long long> cnt(nums.size() + 1);
        long long res = 0;

        cnt[0] = 0;
        for (int i = 1; i <= nums.size(); i++) 
            cnt[i] = cnt[i-1] + ((nums[i-1] % modulo) == k ? 1 : 0);

        for (int i = 0; i < cnt.size(); i++) {
            res += cnt_map[(cnt[i] + modulo - k) % modulo];
            cnt_map[cnt[i] % modulo]++;
        }

        return res;
    }
};
