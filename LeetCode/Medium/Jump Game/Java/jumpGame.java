// recursive solution dp approach
class Solution {
    public boolean canJump(int[] nums) {
        return recur(nums, 1);
    }

    private boolean recur(int[] nums, int ind) {
        if (ind >= nums.length) return true;
        if (nums[ind-1] > 0) {
            if (nums[ind-1] > nums[ind]) nums[ind] = nums[ind-1]-1;
            return recur(nums, ind+1);
        }
        return false;
    }
}

//greedy approach
class Solution {
    public boolean canJump(int[] nums) {
        int jump = nums[0];
        for (int i=1; i<nums.length; i++) {
            if (jump <= 0) return false; 
            jump = Math.max(jump-1, nums[i]);
        }
        return jump >= 0;
    }
}
