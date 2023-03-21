class Solution {
    public long zeroFilledSubarray(int[] nums) {
        int size = nums.length;
        long res = 0;

        for (int i=0; i<size;) {
            if (nums[i] == 0) {
                int k = i;
                long zeros = 0;
                while (k < size && nums[k] == 0) {
                    zeros++;
                    k++;
                }
                res += zeros*(zeros+1)/2;
                i = k;
            } else {
                i++;
            }
        }
        return res;
    }
}
