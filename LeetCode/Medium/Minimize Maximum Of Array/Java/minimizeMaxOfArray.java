class Solution {
    public int minimizeArrayValue(int[] nums) {
        long res =0, prefixSum = 0;
        for (int i=0; i<nums.length; i++) {
            prefixSum += nums[i];
            res = Math.max(res, (prefixSum+i)/(i+1));
        }
        return (int)res;
    }
}
