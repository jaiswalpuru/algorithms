class Solution {
    public long countSubarrays(int[] nums, int minK, int maxK) {
        int size = nums.length;
        int left = -1;
        int minPos = -1;
        int maxPos = -1;
        long subarrays = 0;

        for (int i=0; i<size; i++) {
            if (nums[i]<minK || nums[i]>maxK) left = i;
            if (nums[i] == minK) minPos = i;
            if (nums[i] == maxK) maxPos = i;
            subarrays += Math.max(0, Math.min(minPos, maxPos)-left);
        }

        return subarrays;
    }
}
