class Solution {
    public int maxSubArray(int[] nums) {
        int size = nums.length;
        int sum = (int) -1e5;
        int temp = (int) -1e5;
        for (int i=0; i<size; i++) {
            temp = Math.max(temp+nums[i], nums[i]);
            sum = Math.max(sum, temp);
        }

        return sum;
    }
}
