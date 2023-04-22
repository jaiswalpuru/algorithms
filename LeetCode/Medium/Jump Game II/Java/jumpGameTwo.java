class Solution {
    public int jump(int[] nums) {
        int curr = 0, minJumps = 0, farthestJump = 0;
        int size = nums.length;

        for (int i=0; i<size-1; i++) {
            farthestJump = Math.max(farthestJump, i+nums[i]);
            if (i == curr) {
                minJumps++;
                curr = farthestJump;
            }
        }
        
        return minJumps;
    }
}
