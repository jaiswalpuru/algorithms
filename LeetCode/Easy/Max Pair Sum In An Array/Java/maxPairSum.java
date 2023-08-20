class Solution {
  public int maxSum(int[] nums) {
      int res = -1;
      for (int i=0; i<nums.length; i++) {
          for (int j=i+1; j<nums.length; j++) {
              if (getMaxDigit(nums[i]) == getMaxDigit(nums[j])) {
                  res = Math.max(res, nums[i]+nums[j]);
              }
          }
      }
      return res;
  }

  private int getMaxDigit(int n) {
      int res = 0;
      while(n > 0) {
          res = Math.max(res, n%10);
          n /= 10;
      }
      return res;
  }
}
