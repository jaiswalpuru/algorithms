class Solution {
    public double findMaxAverage(int[] nums, int k) {
        int sum = 0;
        int n = nums.length;
        for (int i=0; i<k; i++) sum += nums[i];
        double avg = (double)(sum)/k;
        if (k == n) return avg;
        int j = k, i = 1;
        sum -= nums[0];
        while (j < n) {
            if (j-i == k) {
                avg = Math.max(avg, (double)(sum)/k);
                sum -= nums[i];
                i++;
            }
            sum += nums[j];
            j++;
        }
        avg = Math.max(avg, (double)(sum)/k);
        return avg;
    }
}
