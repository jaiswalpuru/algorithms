class Solution {
    public int[] rearrangeArray(int[] nums) {
        double median = 0.0;
        Arrays.sort(nums);

        if (nums.length%2 == 0) {
            median = (double)(nums[nums.length/2] + nums[nums.length/2-1])/2;
        }else {
            median = nums[nums.length/2];
        }

        int[] res = new int[nums.length];
        int evenIndex = 0;
        int oddIndex = 1;
        for (int num : nums) {
            if ((double)(num) >= median) { 
                res[evenIndex] = num;
                evenIndex += 2;
            }else {
                res[oddIndex] = num;
                oddIndex += 2;
            }
        }
        
        return res;
    }
}
