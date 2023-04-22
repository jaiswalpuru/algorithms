class Solution {
    public int rob(int[] nums) {
        int size = nums.length;
        int[] memo = new int[size];
        Arrays.fill(memo, -1);
        return Math.max(recur(nums, 0, memo), recur(nums, 1, memo));
    }

    private int recur(int[] nums, int ind, int[] memo) {
        if (ind >= nums.length) return 0;
        if (memo[ind] != -1) return memo[ind];
        int dntTake = recur(nums, ind+1, memo);
        int take = nums[ind] + recur(nums, ind+2, memo);
        memo[ind] = Math.max(take, dntTake);
        return memo[ind];
    }
}
