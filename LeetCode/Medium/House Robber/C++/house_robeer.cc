class Solution {
public:
    int rob(vector<int>& nums) {
        vector<int> memo(nums.size(), -1);
        return max(recur(nums, 0, memo), recur(nums, 1, memo));
    }

    int recur(vector<int> nums, int start, vector<int>& memo) {
        if (start >= nums.size()) return 0;
        if (memo[start] != -1) return memo[start];
        int take = nums[start] + recur(nums, start+2, memo);
        int dnt_take = recur(nums, start+1, memo);
        return memo[start] = max(take, dnt_take);
    }
};
