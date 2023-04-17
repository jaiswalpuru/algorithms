class Solution {
    public int maxProduct(int[] nums) {
        int maxSoFar = nums[0];
        int minSoFar = nums[0];
        int size = nums.length;
        int res = maxSoFar;
        for (int i=1; i<size; i++) {
            int temp = Math.max(nums[i], Math.max(nums[i]*maxSoFar, nums[i]*minSoFar));
            minSoFar = Math.min(nums[i], Math.min(nums[i]*maxSoFar, nums[i]*minSoFar));
            maxSoFar = temp;
            res = Math.max(res, maxSoFar);
        }
        return res;
    }
}
