class Solution {
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        int n = nums.length;
        int i = 0;
        int j = 0;
        int val = 1;
        int cnt = 0;
        while (i < n && j < n ) {
            val *= nums[j];
            while(val >= k && i < n) {
                val /= nums[i];
                i++;
            }
            if (i < n) cnt += j-i+1;
            j++;
        }
        return cnt;
    }
}
