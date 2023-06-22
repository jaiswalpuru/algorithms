class Solution {
    public int minStartValue(int[] nums) {
        int low = 1;
        int high = (int)1e4+1;
        while (low < high) {
            int mid = low + (high-low)/2;
            int sum = mid;
            boolean f = false;
            for (int j=0; j<nums.length; j++) {
                sum += nums[j];
                if (sum <= 0) {
                    f = true;
                    break;
                }
            }
            if (!f) {
                high = mid;    
            } else {
                low = mid+1;
            }
        }
        return low;
    }
}
