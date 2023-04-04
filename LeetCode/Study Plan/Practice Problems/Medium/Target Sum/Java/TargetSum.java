//------------------method one---------------------
class Solution {
    public int findTargetSumWays(int[] nums, int target) {
        return backtrack(nums, 0, target, 0);
    }

    private int backtrack(int[] nums, int ind, int target, int sum) {
        if (ind == nums.length) {
            if (sum == target) return 1;
            return 0;
        }

        int take = backtrack(nums, ind+1, target, sum+nums[ind]);
        int dntTake = backtrack(nums, ind+1, target, sum-nums[ind]);
        return take+dntTake;
    }
}
//------------------method one---------------------

//------------------method two---------------------
class Solution {
    public int findTargetSumWays(int[] nums, int target) {
        int n = nums.length;
        int total = 0;
        for (int val : nums) total += val;

        int[][] memo = new int[n][2*total+1];
        for (int[] row : memo) Arrays.fill(row, Integer.MIN_VALUE);

        return recur(nums, memo, target, 0, 0, total);
    }

    private int recur(int[] nums, int[][] memo, int target, int sum, int ind, int total) {
        if (ind ==  nums.length) {
            if (sum == target) return 1;
            return 0;
        }

        if (memo[ind][sum+total] != Integer.MIN_VALUE) return memo[ind][sum+total];
        int take = recur(nums, memo, target, sum+nums[ind], ind+1, total);
        int dntTake = recur(nums, memo, target, sum-nums[ind], ind+1, total);
        memo[ind][sum+total] = take + dntTake;
        return memo[ind][sum+total];
    }
}
//------------------method two---------------------
