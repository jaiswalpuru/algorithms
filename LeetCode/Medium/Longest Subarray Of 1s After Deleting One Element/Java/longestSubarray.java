class Solution {
    public int longestSubarray(int[] nums) {
        int zero = 0;
        int maxLength = 0;
        int i = 0;
        int k = 0;
        int currLen = 0;
        for (i=0; i<nums.length; i++) {
            if (nums[i] == 0) zero++;
            else currLen++;
            if (zero > 1) {
                maxLength = Math.max(maxLength, i-k-1);
                while (zero > 1) {
                    if (nums[k] == 0) zero--;
                    k++;
                }
            }
        }
        maxLength = Math.max(maxLength, i-k-1);
        return maxLength;
    }
}
