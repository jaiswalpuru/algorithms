class Solution {
    public boolean canJump(int[] nums) {
        int maxJump = nums[0];
        int size = nums.length;
        for (int i=1; i<size; i++) {
            if (maxJump <= 0) return false;
            maxJump = Math.max(maxJump-1, nums[i]);
        }
        return true;
    }
}
