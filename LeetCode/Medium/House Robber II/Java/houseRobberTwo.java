class Solution {
    public int rob(int[] nums) {
        int size = nums.length;
        if (size == 1) return nums[0];
        int[] memoOne = new int[size];
        int[] memoTwo = new int[size];
        Arrays.fill(memoOne, -1);
        Arrays.fill(memoTwo, -1);
        return Math.max(recur(Arrays.copyOfRange(nums, 1, nums.length), 0, memoOne), 
        recur(Arrays.copyOfRange(nums, 0, nums.length-1), 0, memoTwo));
    }

    private int recur(int[] nums, int ind, int[] memo) {
        if (ind >= nums.length) return 0;
        if (memo[ind] != -1) return memo[ind];

        int take = nums[ind] + recur(nums, ind+2, memo);
        int dntTake = recur(nums, ind+1, memo);
        memo[ind] = Math.max(take, dntTake);
        return memo[ind];
    }
}
