class Solution {
public:
    int countMaxOrSubsets(vector<int>& nums) {
        int max_or_val = 0;
        for (int num : nums) max_or_val |= num;
        vector<vector<int>> memo(nums.size(), vector<int>(max_or_val + 1, -1));

        return recur(nums, max_or_val, 0, 0, memo);
    }

    int recur(vector<int> nums, int max_or_val, int ind, int curr_or, vector<vector<int>>& memo) {
        if (ind == nums.size()) {
            return (curr_or == max_or_val) ? 1 : 0;
        }

        if (memo[ind][curr_or] != -1) return memo[ind][curr_or];

        int dnt_take = recur(nums, max_or_val, ind + 1, curr_or, memo);
        int take = recur(nums, max_or_val, ind + 1, curr_or | nums[ind], memo);
        
        return memo[ind][curr_or] = take + dnt_take;
    } 
};
