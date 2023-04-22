class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        int size = nums.length;
        int res = Integer.MAX_VALUE;
        int sum = 0;
        int k = 0;

        for (int i=0; i<size; i++) {
            sum += nums[i];
            while(sum >= target) {
                res = Math.min(res, i-k+1);
                sum -= nums[k];
                k++;
            }
        }

        return res == Integer.MAX_VALUE ? 0 : res;
    }
}
