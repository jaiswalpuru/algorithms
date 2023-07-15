class Solution {
    public int jump(int[] nums) {
        int res = 0;
        int currFar = 0, currEnd = 0;
        int n = nums.length;
        for (int i=0; i<n-1; i++) {
            currFar = Math.max(currFar, i+nums[i]);
            if (i == currEnd) {
                System.out.println(i + " " + currFar);
                res++;
                currEnd = currFar;
            }
        }
        return res;
    }
}
