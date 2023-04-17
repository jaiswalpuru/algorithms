class Solution {
    public int lengthOfLIS(int[] nums) {
        int size = nums.length;
        int[][] memo = new int[size][size+1];
        for (int i=0; i<size; i++) Arrays.fill(memo[i], -1);
        return recur(nums, 0, -1, memo);
    }

    private int recur(int[] nums, int ind, int prev, int[][] memo) {
        if (ind == nums.length) return 0;
        if (memo[ind][prev+1] != -1) return memo[ind][prev+1];

        int dntTake = recur(nums, ind+1, prev, memo);
        int take = 0;
        if (prev == -1 || (ind > prev && nums[prev] < nums[ind])) {
            take = 1+recur(nums, ind+1, ind, memo);
        }
        memo[ind][prev+1] = Math.max(dntTake, take);
        return memo[ind][prev+1];
    }
}
