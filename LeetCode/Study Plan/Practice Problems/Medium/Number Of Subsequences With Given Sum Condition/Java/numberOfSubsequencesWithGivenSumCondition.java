class Solution {
    int mod = (int)1e9+7;
    public int numSubseq(int[] nums, int target) {
        Arrays.sort(nums); 
        int size = nums.length;
        int[] powers2 = new int[size];
        
        powers2[0] = 1;
        for (int i=1; i<size; i++)
            powers2[i] = (powers2[i-1]*2)%mod;
        
        int left = 0, right = size-1;
        int res = 0;
        while(left <= right) {
            if (nums[left] + nums[right] <= target) {
                res = (res + powers2[right-left])%mod;
                left++;
            }else {
                right--;
            }
        }

        return res;
    }
}
