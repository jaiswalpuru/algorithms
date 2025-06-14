class Solution {
public:
    vector<vector<int>> res;
    vector<vector<int>> threeSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size(); i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            two_sum(nums, nums[i], i + 1);
        }
        return res;
    }

    void two_sum(vector<int>& nums, int curr, int i) {
        int l = i;
        int r = nums.size() - 1;
        vector<int> temp;
        while (l < r) {
            int sum = nums[l] + nums[r] + curr;
            if (sum == 0) {
                res.push_back({curr, nums[l], nums[r]});
                l++;
                r--;
                while (l < r && nums[l] == nums[l - 1]) l++;
            } else if (sum > 0) {
                r--;
            } else {
                l++;
            }
        }
    }
};
