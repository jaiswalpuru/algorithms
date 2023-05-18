class Solution {
    public int maxOperations(int[] nums, int k) {
        Arrays.sort(nums);
        int i=0, j=nums.length-1;
        int cnt = 0;
        while(i < j) {
            int sum = nums[i]+nums[j];
            if (sum == k) {
                cnt++;
                i++;
                j--;
            }
            else if (sum > k) j--;
            else i++;
        }
        return cnt;
    }
}
