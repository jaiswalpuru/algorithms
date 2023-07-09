class Solution {
    public int removeDuplicates(int[] nums) {
        int i=0;
        int j=0;
        int n = nums.length;
        while (j < n) {
            int k=j+1;
            while(k < n && nums[j] == nums[k]) k++;
            nums[i] = nums[j];
            i++;
            j=k;
        }
        return i;
    }
}
