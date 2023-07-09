class Solution {
    public int removeElement(int[] nums, int val) {
        int k=nums.length-1;
        int i=0;
        if (nums.length == 0) return 0;
        while (i<k) {
            if (nums[i] == val){
                nums[i] = nums[k];
                k--;
            }else {
                i++;
            }
        }
        return nums[k] == val ? k : k+1;
    }
}
