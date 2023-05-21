class Solution {
    public int pivotIndex(int[] nums) {
        int n = nums.length;
        int[] pre = new int[n+1];
        pre[0] = 0;
        for (int i=1; i<=n; i++)
            pre[i] = pre[i-1] + nums[i-1];

        for (int i=1; i<=n; i++) {
            int left = pre[i-1];
            int right = pre[n]-pre[i-1]-nums[i-1];
            if (left == right) return i-1;
        }
        return -1;
    }
}
