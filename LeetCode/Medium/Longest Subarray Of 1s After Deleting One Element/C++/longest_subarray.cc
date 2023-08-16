class Solution {
public:
    int longestSubarray(vector<int>& nums) {
        int z=0, j=0;
        int max_len = 0;
        int i=0;
        for (i=0; i<nums.size(); i++) {
            if (nums[i] == 0) z++;
            if (z > 1) {
                max_len = max(max_len, i-j-1);
                while (z > 1) {
                    if (nums[j] == 0) {
                        z--;
                    }
                    j++;
                }
            }
        }
        max_len = max(max_len, i-j-1);
        return max_len;
    }
};
