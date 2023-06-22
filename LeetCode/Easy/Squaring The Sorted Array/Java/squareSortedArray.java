class Solution {
    public int[] sortedSquares(int[] nums) {
        int n = nums.length;
        int l=0;
        int r = n-1;
        int[] res = new int[n];
        int val = -1;

        for (int i=n-1; i>=0; i--) {
            if (Math.abs(nums[l]) < Math.abs(nums[r])) val = nums[r--];
            else val = nums[l++];
            res[i] = val*val;
        }

        return res;
    }
}
