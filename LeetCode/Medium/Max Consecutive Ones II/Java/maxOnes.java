class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        int res=0, l=0, r=0;
        int zero = 0;
        while(r < nums.length) {
            if (nums[r] == 0) zero++;
            while(zero > 1) {
                if (nums[l]==0)zero--;
                l++;
            }
            res = Math.max(res, r-l+1);
            r++;
        }
        return res;
    }
}
