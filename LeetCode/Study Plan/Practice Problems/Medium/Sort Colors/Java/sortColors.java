class Solution {
    public void sortColors(int[] nums) {
        int n = nums.length;
        int f = 0, l = n-1;
        int curr = 0;
        int p0 = 0, p2 = n-1;
        while(curr <= p2) {
            if (nums[curr] == 0) {
                int temp = nums[curr];
                nums[curr] = nums[p0];
                nums[p0] = temp;
                curr++;
                p0++;
            }else if (nums[curr] == 2) {
                int temp = nums[curr];
                nums[curr] = nums[p2];
                nums[p2] = temp;
                p2--;
            }else {
                curr++;
            }
        }
    }
}
