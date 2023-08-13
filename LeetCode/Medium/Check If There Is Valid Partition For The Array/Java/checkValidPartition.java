class Solution {
    Map<Integer, Boolean> memo;
    public boolean validPartition(int[] nums) {
        memo = new HashMap<>();
        memo.put(-1, true);
        return recur(nums, nums.length-1);
    }

    private boolean recur(int[] nums, int i) {
        if (memo.containsKey(i)) return memo.get(i);
        boolean res = false;
        if (i > 0 && nums[i] == nums[i-1]) res |= recur(nums, i-2);
        if (i > 1 && nums[i] == nums[i-1] && nums[i-1] == nums[i-2]) res |= recur(nums, i-3);
        if (i > 1 && nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1) res |= recur(nums, i-3);
        memo.put(i, res);
        return res;
    }
}
