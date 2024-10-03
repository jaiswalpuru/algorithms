class Solution {
public:
    int minSubarray(vector<int>& nums, int p) {
        int target = (accumulate(nums.begin(), nums.end(), 0LL) % p);
        if (target == 0) return 0;
        int curr_sum = 0;
        int min_len = nums.size();
        unordered_map<int, int> umap;

        umap[0] = -1;
        for (int i = 0; i < nums.size(); i++) {
            curr_sum = (curr_sum + nums[i]) % p;
            int need = (curr_sum - target + p) % p;
            if (umap.find(need) != umap.end()) min_len = min(min_len, i - umap[need]);
            umap[curr_sum] = i;
        }

        return min_len == nums.size() ? -1 : min_len;
    }
};
