class Solution {
    public int[] getAverages(int[] nums, int k) {
        int n = nums.length;
        int[] avg = new int[n];
        Arrays.fill(avg, -1);
        long[] pre = new long[n];
        pre[0] = nums[0];
        for (int i=1; i<n; i++) pre[i] = pre[i-1]+nums[i];
        for (int c=k; c+k<n && c-k>=0; c++) {
            if (c-k == 0) {
                avg[c] = (int)(pre[c+k]/(k*2+1));
            } else {
                avg[c] = (int)((pre[c+k]-pre[c-k-1])/(k*2+1));
            }
        }
        return avg;
    }
}
