class Solution {
    public int removeDuplicates(int[] nums) {
        int i=0;
        int j=0;
        int n = nums.length;
        while (j < n) {
            int k = j+1;
            while(k < n && nums[k] == nums[j]) k++;
            if (k-j >= 2) {
                nums[i] = nums[j];
                nums[i+1] = nums[j];
                i += 2;
            } else {
                nums[i] = nums[j];
                i++;
            }
            j = k;
        }
        return i;
    }
}
