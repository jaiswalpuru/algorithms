class Solution {
public:
    int longestSubarray(vector<int>& nums) {
        int res = 0;
        int size = 0;
        int max_val = 0;
        
        for (int i = 0; i < nums.size(); i++) max_val = max(max_val, nums[i]);
        for (int i = 0; i < nums.size(); i++) {
            if (max_val == nums[i]) size++;
            else size = 0;
            res = max(res, size);
        }
        
        return res;
    }
};
